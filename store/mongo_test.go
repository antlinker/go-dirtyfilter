package store_test

import (
	"github.com/antlinker/go-dirtyfilter/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("敏感词MongoDB存储测试", func() {
	var (
		mgoStore *store.MongoStore
	)
	BeforeEach(func() {
		s, err := store.NewMongoStore(store.MongoConfig{
			URL: "mongodb://admin:123456@192.168.33.70:27017",
			DB:  "sample",
		})
		if err != nil {
			Fail(err.Error())
			return
		}
		mgoStore = s
		err = s.Write("共产党")
		if err != nil {
			Fail(err.Error())
			return
		}
	})
	It("Read Test", func() {
		for v := range mgoStore.Read() {
			Expect(v).To(Equal("共产党"))
		}
	})
	It("ReadAll Test", func() {
		result, err := mgoStore.ReadAll()
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(result).To(Equal([]string{"共产党"}))
	})
	It("Version Test", func() {
		Expect(mgoStore.Version()).To(Equal(uint64(1)))
		err := mgoStore.Write("党")
		if err != nil {
			Fail(err.Error())
			return
		}
		Expect(mgoStore.Version()).To(Equal(uint64(2)))
		err = mgoStore.Remove("党")
		if err != nil {
			Fail(err.Error())
			return
		}
	})
	AfterEach(func() {
		err := mgoStore.Remove("共产党")
		if err != nil {
			Fail(err.Error())
			return
		}
	})
})
