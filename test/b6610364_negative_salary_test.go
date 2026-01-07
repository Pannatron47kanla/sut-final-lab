package entity

import (
	"finallab/entity"
	"testing"

	"github.com/onsi/gomega"
)

func TestSalary(t *testing.T) {
	g := gomega.NewWithT(t)
	var employees = entity.Employees{
		Name:         "jonh nathan",
		Salary:       1,
		EmployeeCode: "HR-1024",
	}

	err := employees.Validate()
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("Salary must be between 15000 and 200000"))
}
