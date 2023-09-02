package jwt

import "time"

var accessSecretKey = []byte("1!@#$fkdaowekvmzxce")
var refreshSecretKey = []byte("9f,a-023!@fnfu349)*(&0!adv124")
var accessLifeTime = time.Nanosecond * 60 * 30
var refreshLifeTime = time.Nanosecond * 60 * 60 * 3
