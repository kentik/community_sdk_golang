package cloudexportstub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
)

const cloudExportNotFound = -1

// CloudExportRepo is local storage fro CloudExport data
type CloudExportRepo struct {
	fileName string
	lock     sync.RWMutex
	items    []V202101beta1CloudExport
}

func NewCloudExportRepo(fileName string) *CloudExportRepo {
	r := &CloudExportRepo{
		fileName: fileName,
		items:    make([]V202101beta1CloudExport, 0),
	}
	r.load()
	return r
}

func (r *CloudExportRepo) Get(id string) *V202101beta1CloudExport {
	r.lock.RLock()
	defer r.lock.RUnlock()

	index := r.findByID(id)
	if index == cloudExportNotFound {
		return nil
	}

	copyExport := r.items[index]
	return &copyExport
}

func (r *CloudExportRepo) List() []V202101beta1CloudExport {
	r.lock.RLock()
	defer r.lock.RUnlock()

	copyitems := make([]V202101beta1CloudExport, len(r.items), len(r.items))
	copy(copyitems, r.items)
	return copyitems
}

func (r *CloudExportRepo) Create(e V202101beta1CloudExport) (*V202101beta1CloudExport, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.findByName(e.Name) != cloudExportNotFound {
		return nil, fmt.Errorf("cloud export %q already exists", e.Name)
	}

	newExport := e

	// cloudexport service assigns ApiRoot and FlowDest if it is not specified by user
	if e.ApiRoot == nil {
		newExport.ApiRoot = stringPtr("http://localhost:8080/api")
	}
	if e.FlowDest == nil {
		newExport.FlowDest = stringPtr("http://localhost:8080/flow")
	}

	newExport.Id = r.allocateNewID()
	newExport.CurrentStatus = &CloudExportv202101beta1Status{
		Status:               "OK",
		ErrorMessage:         "No errors",
		FlowFound:            true,
		ApiAccess:            true,
		StorageAccountAccess: true,
	}

	r.items = append(r.items, newExport)
	r.save()
	return &newExport, nil
}

func (r *CloudExportRepo) Update(e V202101beta1CloudExport) (*V202101beta1CloudExport, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	exportIndex := r.findByID(e.Id)
	if exportIndex == cloudExportNotFound {
		return nil, fmt.Errorf("cloud export of id %q doesn't exists", e.Id)
	}
	r.items[exportIndex] = e
	r.save()
	return &r.items[exportIndex], nil
}

func (r *CloudExportRepo) Delete(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	exportIndex := r.findByID(id)
	if exportIndex == cloudExportNotFound {
		return fmt.Errorf("cloud export of id %q doesn't exists", id)
	}
	r.items = append(r.items[:exportIndex], r.items[exportIndex+1:]...)
	r.save()
	return nil
}

func (r *CloudExportRepo) findByID(id string) int {
	// r.lock should be already in locked state at this point
	for i, item := range r.items {
		if item.Id == id {
			return i
		}
	}
	return cloudExportNotFound
}

func (r *CloudExportRepo) findByName(name string) int {
	// r.lock should be already in locked state at this point
	for i, item := range r.items {
		if item.Name == name {
			return i
		}
	}
	return cloudExportNotFound
}

func (r *CloudExportRepo) allocateNewID() string {
	// r.lock should be already in locked state at this point
	var id int

	for _, item := range r.items {
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

func stringPtr(s string) *string {
	return &s
}

func (r *CloudExportRepo) load() {
	// r.lock should be already in locked state at this point
	bytes, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		log.Printf("CloudExportRepo storage %q not found. Starting with fresh one", r.fileName)
		return // nothing to load is ok
	}

	err = json.Unmarshal(bytes, &r.items)
	if err != nil {
		log.Printf("CloudExportRepo unmarshal from %q failed: %v", r.fileName, err)
		panic(err) // loading invalid data is not ok
	}

	log.Printf("CloudExportRepo successfuly loaded from %q, num exports: %d", r.fileName, len(r.items))
}

func (r *CloudExportRepo) save() {
	// r.lock should be already in locked state at this point
	bytes, err := json.MarshalIndent(r.items, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(r.fileName, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
