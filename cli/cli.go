package cli

import "github.com/AlecAivazis/survey/v2"
import "math"
import "go-gabut/loop"
import "go-gabut/pattern"
import "go-gabut/console"

func factorial(value int) int {
    if value == 1 {
        return 1
    } else {
        result := value * factorial(value - 1)
        return result
    }
}

func exponen(xx, yy int) int {
    return int(math.Pow(float64(xx), float64(yy)))
}

func getType() string {
    var typeMath string
    typeMathPrompt := &survey.Input{
        Message: "Enter Math Type",
    }
    
    err := survey.Ask([]*survey.Question{{
        Name: "typeMath",
        Prompt: typeMathPrompt,
        Validate: survey.Required,
    }}, &typeMath)
    if err != nil {
		return ""
    }
    return typeMath
}

func getExponen() int {
    var exp int
    expPrompt := &survey.Input{
        Message: "Enter Rank",
    }
    
    err := survey.Ask([]*survey.Question{{
        Name: "exp",
        Prompt: expPrompt,
        Validate: survey.Required,
    }}, &exp)
    if err != nil {
		return 0
    }
    return exp
}

func GetMath() {
    console.Log("WELCOME TO MATH TOOLS")
	var num int
	numPrompt := &survey.Input{
		Message: "Enter a Number",
	}
	err := survey.Ask([]*survey.Question{{
		Name:    "num",
		Prompt:  numPrompt,
		Validate: survey.Required,
	}}, &num)
	
	if err != nil {
		return
	}
	
	var typeMath = getType()
	switch typeMath {
	case "factorial":
	    console.Log(factorial(num))
	    break
	case "exponen":
	    exp := getExponen()
	    console.Log(exponen(num, exp))
	case "cordinate":
	    loop.Cordinate(num)
	case "diamond":
	    pattern.Diamond(num)
	}
} 