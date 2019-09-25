package collatzconjecture
import "errors"

// This should return the collatzConjecture output of given number
func CollatzConjecture(input int) (int,error){
   counter:=0
   if input<=0{
	 return input,errors.New("Invalid value provided")
   }else{
   for input>1{
     if input%2==0{
        input=input/2
     } else{
      input=(3*input)+1
    }
   counter=counter+1
   }
   return counter,nil
   }
   }

