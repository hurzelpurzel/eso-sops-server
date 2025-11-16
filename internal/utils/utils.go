package utils

import (
	"fmt"
	"os"

)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}




func GetEnvOrFail (env string) (string, error){
	if value, exists := os.LookupEnv(env); exists {
        return value, nil 
    }else{
		return "", fmt.Errorf("failed to load env var %s", env)
	}
}