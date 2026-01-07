package entity

import (
	"finallab/entity"
	"testing"

	"github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := gomega.NewWithT(t)
	var employees = entity.Employees{
		Name:         "jonh nathan",
		Salary:       20000,
		EmployeeCode: "HR-1024",
	}

	err := employees.Validate()
	g.Expect(err).To(gomega.BeNil())
}
