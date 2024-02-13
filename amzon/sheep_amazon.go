package main

import "fmt"

/**
 *
 टिम्मी भेड़ ने एक ऐसी रणनीति बनाई है जो उसे जल्दी सो जाने में मदद करती है।
 सबसे पहले, वह एक संख्या N चुनती है। फिर वह N, 2 × N, 3 × N, इत्यादि नाम रखना शुरू करती है।
  जब भी वह किसी संख्या का नाम बताती है तो उस संख्या के सभी अंकों के बारे में सोचती है।
  वह किस अंक (0, 1, 2, 3, 4, 5, 6, 7, 8, और 9) पर नज़र रखती है
  उसने अब तक कम से कम एक बार अपने द्वारा नामित किसी भी नंबर का हिस्सा देखा है।
  एक बार जब वह दस अंकों में से प्रत्येक को कम से कम एक बार देख लेगी, तो वह सो जाएगी।
टिम्मी को N से शुरू करना चाहिए और हमेशा i × N के बाद सीधे (i + 1) × N नाम रखना चाहिए।
उदाहरण के लिए, मान लीजिए कि टिम्मी ने N = 1692 चुना है। वह इस प्रकार गिनती करेगी:

एन = 1692. अब उसने अंक 1, 2, 6 और 9 देखे हैं।
2N = 3384. अब उसने अंक 1, 2, 3, 4, 6, 8 और 9 देखे हैं।
3एन = 5076। अब उसने सभी दस अंक देख लिए हैं और सो जाती है।
सोने से पहले वह आखिरी नंबर कौन सा बताएगी? यदि वह हमेशा के लिए गिनती करेगी, तो इसके बजाय INSOMNIA प्रिंट करें।
*/
/*
n=1692
dp[]
a =1692
i=1
ans=1692
dp[2]=true,  a =169
dp[9]=true ,a =16
dp[6]=true ,a =1
dp[1]=true ,a =0

i=2
a=3384
ans=3384

dp[4]=true, a =338
dp[8]=true ,a =33
dp[3]=true ,a =3
dp[1]=true ,a =0


 1 ->  all digites
 2 -> 2,4,6,8,10, ->all digites
 3 -> 3,6,9,12,15 - > all digites
 99 -> 1,8,6 ,5,9,7

*/
func findLastNum(n int) int {
	uniqueDigites := make(map[int]bool)
	if n == 0 {
		return -1
	}
	i := 1
	ans := -1
	for true {
		nextNumber := i * n
		ans = nextNumber
		for nextNumber > 0 { // 1692/10 = 169, 169/10= 16
			rem := nextNumber % 10       // take last digite from end
			nextNumber = nextNumber / 10 // remaining digites after removing last one
			if !uniqueDigites[rem] {     // check if uniqueDigites don't have digit then put it into map
				uniqueDigites[rem] = true
			}
			if len(uniqueDigites) == 10 { // check if uniqueDigites size is 10 then we got all digites from
				return ans
			}
		}
		i++
	}
	return ans
}

func main() {
	fmt.Println(findLastNum(1692))
	/*for i := 0; i < 1000; i++ {
		fmt.Println(findLastNum(i))
	}*/
}
