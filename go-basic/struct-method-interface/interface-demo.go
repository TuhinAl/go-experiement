package interface_test

import "fmt"

func main9() {
	/*	name := MyString("alauddin tuhin")
		var vowelsFinder VowelsFinder
		vowelsFinder = name // possible since MyString implements VowelsFinder
		fmt.Printf("Vowels are %c", vowelsFinder.FindVowels())*/

	employees := sliceOfSalaryCalculator()
	totalExpense(employees)
	//fmt.Printf("Total Expenses are %d", total)
}

type VowelsFinder interface {
	FindVowels() []rune // rune is alias of int 32
}

type MyString string

//MyString implements VowelsFinder

func (myString MyString) FindVowels() []rune {
	var vowels []rune

	for _, vowel := range myString {
		if vowel == 'a' || vowel == 'e' || vowel == 'i' || vowel == 'o' || vowel == 'u' {
			vowels = append(vowels, vowel)

		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	employeeId  int
	basicSalary int
	pf          int
}

type Contract struct {
	employeeId  int
	basicSalary int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (permanent Permanent) CalculateSalary() int {
	//grantPF := permanent.basicSalary * .07;
	return permanent.basicSalary + permanent.pf
}

func (contract Contract) CalculateSalary() int {
	return contract.basicSalary
}

func (freelancer Freelancer) CalculateSalary() int {
	return freelancer.ratePerHour * freelancer.totalHours * 22
}

func totalExpense(salaryCalculator []SalaryCalculator) {
	totalExpense := 0
	for _, salary := range salaryCalculator {
		totalExpense += salary.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month ৳%d", totalExpense)
}

func sliceOfSalaryCalculator() []SalaryCalculator {
	selim := Permanent{100, 50000, 3500}
	josim := Permanent{101, 45000, 2700}
	bapparaj := Contract{102, 70000}
	tanvir := Freelancer{103, 340, 8}
	freelancer2 := Freelancer{104, 220, 9}
	return []SalaryCalculator{selim, josim, bapparaj, tanvir, freelancer2}
}
