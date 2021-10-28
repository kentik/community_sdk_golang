package syntheticsstub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

const itemNotFound = -1

// SyntheticsRepo is local storage for Synthetics data
type SyntheticsRepo struct {
	fileName        string
	lock            sync.RWMutex
	tests           []V202101beta1Test
	agents          []V202101beta1Agent
	health          []V202101beta1TestHealth
	lookups         []V202101beta1TracerouteLookup
	traceRoutes     []V202101beta1TracerouteResult
	traceRoutesInfo []V202101beta1TracerouteInfo
}

// syntheticsStorage is internally used for marshall/unmarshall the data as JSON
type syntheticsStorage struct {
	Tests           []V202101beta1Test
	Agents          []V202101beta1Agent
	Health          []V202101beta1TestHealth
	Lookups         []V202101beta1TracerouteLookup
	TraceRoutes     []V202101beta1TracerouteResult
	TraceRoutesInfo []V202101beta1TracerouteInfo
}

func NewSyntheticsRepo(fileName string) *SyntheticsRepo {
	r := &SyntheticsRepo{
		fileName: fileName,
		tests:    make([]V202101beta1Test, 0),
		agents:   make([]V202101beta1Agent, 0),
	}
	r.load()
	return r
}

func (r *SyntheticsRepo) GetTest(id string) *V202101beta1Test {
	r.lock.RLock()
	defer r.lock.RUnlock()

	index := r.findTestByID(id)
	if index == itemNotFound {
		return nil
	}

	copyItem := r.tests[index]
	return &copyItem
}

func (r *SyntheticsRepo) ListTests() []V202101beta1Test {
	r.lock.RLock()
	defer r.lock.RUnlock()

	copyItems := make([]V202101beta1Test, len(r.tests), len(r.tests))
	copy(copyItems, r.tests)
	return copyItems
}

// NOTE: multiple tests with same name are allowed
func (r *SyntheticsRepo) CreateTest(t V202101beta1Test) (*V202101beta1Test, error) {
	// NOTE: add input "t" validation here if it turns out to be useful

	r.lock.Lock()
	defer r.lock.Unlock()

	newTest := t
	newTest.Id = r.allocateNewTestID()
	newTest.Cdate = time.Now()

	r.tests = append(r.tests, newTest)
	r.save()
	return &newTest, nil
}

// NOTE: in live server PATCH comes with update mask; here we just update the entire test object for simplicity
func (r *SyntheticsRepo) PatchTest(t V202101beta1Test) (*V202101beta1Test, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := r.findTestByID(t.Id)
	if index == itemNotFound {
		return nil, fmt.Errorf("test of id %q doesn't exists", t.Id)
	}
	r.tests[index] = t
	r.save()
	return &t, nil
}

func (r *SyntheticsRepo) DeleteTest(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := r.findTestByID(id)
	if index == itemNotFound {
		return fmt.Errorf("test of id %q doesn't exists", id)
	}
	r.tests = append(r.tests[:index], r.tests[index+1:]...)
	r.save()
	return nil
}

func (r *SyntheticsRepo) UpdateTestStatus(id string, status V202101beta1TestStatus) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := r.findTestByID(id)
	if index == itemNotFound {
		return fmt.Errorf("test of id %q doesn't exists", id)
	}

	r.tests[index].Status = status
	return nil
}

func (r *SyntheticsRepo) DeleteAgent(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	itemIndex := r.findAgentByID(id)
	if itemIndex == itemNotFound {
		return fmt.Errorf("agent of id %q doesn't exists", id)
	}
	r.agents = append(r.agents[:itemIndex], r.agents[itemIndex+1:]...)
	r.save()
	return nil
}

func (r *SyntheticsRepo) GetAgent(id string) *V202101beta1Agent {
	r.lock.RLock()
	defer r.lock.RUnlock()

	index := r.findAgentByID(id)
	if index == itemNotFound {
		return nil
	}

	copyItem := r.agents[index]
	return &copyItem
}

// NOTE: in live server PATCH comes with update mask; here we just update the entire agent object for simplicity
func (r *SyntheticsRepo) PatchAgent(a V202101beta1Agent) (*V202101beta1Agent, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := r.findAgentByID(a.Id)
	if index == itemNotFound {
		return nil, fmt.Errorf("agent of id %q doesn't exists", a.Id)
	}
	r.agents[index] = a
	r.save()
	return &r.agents[index], nil
}

func (r *SyntheticsRepo) ListAgents() []V202101beta1Agent {
	r.lock.RLock()
	defer r.lock.RUnlock()

	copyItems := make([]V202101beta1Agent, len(r.agents), len(r.agents))
	copy(copyItems, r.agents)
	return copyItems
}

// NOTE: in live server GetHealthForTests comes with set of filters; here we just return same result everytime for simplicity
func (r *SyntheticsRepo) GetHealthForTests() []V202101beta1TestHealth {
	// no lock needed as this collection is read-only
	copyItems := make([]V202101beta1TestHealth, len(r.health), len(r.health))
	copy(copyItems, r.health)
	return copyItems
}

// TODO: Implement logic behind GetLookups and GetTraceRoutesInfo
func (r *SyntheticsRepo) GetLookups(t V202101beta1GetTraceForTestRequest) V202101beta1TracerouteLookup {
	return V202101beta1TracerouteLookup{}
}

// NOTE: in live server GetTraceRoutes comes with set of filters; here we just return same result everytime for simplicity
func (r *SyntheticsRepo) GetTraceRoutes() []V202101beta1TracerouteResult {
	// no lock needed as this collection is read-only
	copyItems := make([]V202101beta1TracerouteResult, len(r.traceRoutes), len(r.traceRoutes))
	copy(copyItems, r.traceRoutes)
	return copyItems
}

func (r *SyntheticsRepo) GetTraceRoutesInfo(t V202101beta1GetTraceForTestRequest) V202101beta1TracerouteInfo {
	return V202101beta1TracerouteInfo{}
}

func (r *SyntheticsRepo) findTestByID(id string) int {
	// r.lock should be already in locked state at this point
	for i, item := range r.tests {
		if item.Id == id {
			return i
		}
	}
	return itemNotFound
}

func (r *SyntheticsRepo) findAgentByID(id string) int {
	// r.lock should be already in locked state at this point
	for i, item := range r.agents {
		if item.Id == id {
			return i
		}
	}
	return itemNotFound
}

func (r *SyntheticsRepo) allocateNewTestID() string {
	// r.lock should be already in locked state at this point
	var id int

	for _, item := range r.tests {
		itemID, err := strconv.Atoi(item.Id)
		if err != nil {
			itemID = 1000000 // str conversion error, assume some high integer for the id
		}
		if itemID > id {
			id = itemID
		}
	}
	return strconv.FormatInt(int64(id)+1, 10)
}

func (r *SyntheticsRepo) load() {
	// r.lock should be already in locked state at this point
	bytes, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		log.Printf("SyntheticsRepo storage %q not found. Starting with fresh one", r.fileName)
		return // nothing to load is ok
	}

	var storage syntheticsStorage
	err = json.Unmarshal(bytes, &storage)
	if err != nil {
		log.Printf("SyntheticsRepo unmarshal from %q failed: %v", r.fileName, err)
		panic(err) // loading invalid data is not ok
	}

	r.tests = storage.Tests
	r.agents = storage.Agents
	r.health = storage.Health
	r.lookups = storage.Lookups
	r.traceRoutes = storage.TraceRoutes
	r.traceRoutesInfo = storage.TraceRoutesInfo

	log.Printf("SyntheticsRepo successfuly loaded from %q, "+
		"num tests: %d, "+
		"num agents: %d, "+
		"num health: %d, "+
		"num lookups: %d, "+
		"num traceRoutes: %d, "+
		"num traceRoutesInfo: %d",
		r.fileName,
		len(r.tests),
		len(r.agents),
		len(r.health),
		len(r.lookups),
		len(r.traceRoutes),
		len(r.traceRoutesInfo))
}

func (r *SyntheticsRepo) save() {
	// r.lock should be already in locked state at this point
	storage := syntheticsStorage{
		Tests:           r.tests,
		Agents:          r.agents,
		Health:          r.health,
		Lookups:         r.lookups,
		TraceRoutes:     r.traceRoutes,
		TraceRoutesInfo: r.traceRoutesInfo,
	}

	bytes, err := json.MarshalIndent(&storage, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(r.fileName, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
