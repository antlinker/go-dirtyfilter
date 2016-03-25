package store_test

import (
	"github.com/antlinker/go-dirtyfilter/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("敏感词内存存储测试", func() {
	var (
		memStore *store.MemoryStore
	)
	BeforeEach(func() {
		s, err := store.NewMemoryStore(store.MemoryConfig{
			DataSource: []string{"共产党"},
		})
		if err != nil {
			Fail(err.Error())
			return
		}
		memStore = s
	})
	It("Write Test", func() {
		err := memStore.Write("党")
		if err != nil {
			Fail(err.Error())
			return
		}
		result, err := memStore.ReadAll()
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(len(result)).To(Equal(2))
	})
	It("Read Test", func() {
		for v := range memStore.Read() {
			Expect(v).To(Equal("共产党"))
		}
	})
	It("ReadAll Test", func() {
		result, err := memStore.ReadAll()
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(result).To(Equal([]string{"共产党"}))
	})
	It("Remove Test", func() {
		err := memStore.Remove("共产党")
		if err != nil {
			Fail(err.Error())
			return
		}
		result, err := memStore.ReadAll()
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(len(result)).To(Equal(0))
	})
	It("Version Test", func() {
		Expect(memStore.Version()).To(Equal(uint64(0)))
		err := memStore.Write("党")
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(memStore.Version()).To(Equal(uint64(1)))
	})
})
