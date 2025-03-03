package zng_tool

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFloat(t *testing.T) {

	var slice = []float64{
		1,
		1,
		0.27,
		65.88,
		13.56,
		9.49,
		6.78,
		1.35,
		0.67,
	}

	var oldRate float64

	var rate = new(big.Float)
	for _, v := range slice {
		rate = rate.Add(rate, big.NewFloat(v))
		oldRate = oldRate + v
	}
	f, accuracy := rate.Float32()
	s := accuracy.String()
	fmt.Println(f, s, oldRate)

}

type Permission int32

const (
	PersonalPermission   Permission = 1 << iota // 0001 (二进制) -> 1 (十进制)
	EnterprisePermission                        // 0010 (二进制) -> 2 (十进制)
)

// Set
func (p Permission) Set(userType Permission) int32 {
	return int32(p | userType)
}

// 取消权限
func (p Permission) Cancel(userType Permission) int32 {
	return int32(p &^ userType)
}

// 判断权限
func (p Permission) Check(userType Permission) bool {
	return p&userType != 0
}

func (p Permission) All() bool {
	var personal = p&PersonalPermission != 0
	var enterprise = p&EnterprisePermission != 0
	if enterprise && personal {
		return true
	}
	return false
}

// 设置权限
func SetPermission(permission int, userType int) int {
	return permission | userType
}

// 取消权限
func CancelPermission(permission int, userType Permission) int {
	return permission &^ int(userType)
}

// 判断权限
func CheckPermission(permission int, userType Permission) bool {
	return permission&int(userType) != 0
}

// 获取权限十进制值
func GetPermission(permission int) int {
	return permission
}

func TestPermission(t *testing.T) {

	permissionNew := CancelPermission(3, EnterprisePermission)
	permissionNew = CancelPermission(permissionNew, PersonalPermission)
	fmt.Println(permissionNew)
	//将PersonalPermission

}
