package main

import "fmt"

/*
दस्तावेज़ संग्रहण प्रणाली (Google डॉक्स के समान) उपयोगकर्ता और दस्तावेज़ बनाने की क्षमता एक दस्तावेज़ है
सरलता के लिए वर्णों की एक श्रृंखला। दस्तावेज़ बनाने वाला उपयोगकर्ता इसका एकमात्र स्वामी है
दस्तावेज़। अनुमति प्रबंधन. स्वामी के लिए पढ़ने/लिखने की पहुंच आवंटित करने की क्षमता
 अन्य उपयोगकर्ताओं के लिए. केवल सही अनुमति वाले उपयोगकर्ताओं को ही एक्सेस करने में सक्षम होना चाहिए। गर्म और ठंडे स्तर का भंडारण।
  एप्लिकेशन बंद होने पर भी कोल्ड टियर स्टोरेज उपलब्ध रहेगा। स्टोर के लिए बाहरी डेटाबेस का उपयोग न करें


User
Docs
Permission
*/

var UserMap map[int]User

func main() {
	UserMap = make(map[int]User)
	startDocs()

}

func startDocs() {
	docName1 := "doc_1"

	createUser(23432, "sheelendar")
	createUser(12321, "Niki")
	user1 := UserMap[23432]
	user1.CreateDoc(docName1)

	printDocs(*user1.ListDocs[docName1])
	user1.GivenPermission(docName1, 12321, Read)
	printDocs(*user1.ListDocs[docName1])

	user1Doc1 := user1.ListDocs[docName1]
	user1Doc1.AppendInDoc("my name sheelendar")

	printDocs(*user1.ListDocs[docName1])
}

func printDocs(doc Docs) {
	fmt.Print("Msg:  ", doc.Doc, ",  ")
	fmt.Print("doc_name:   ", doc.Name, "  ", doc.CreatedTime, "   ", doc.SharedUsers)
	fmt.Println()
}
func createUser(userId int, userName string) {
	user := User{UserID: userId, UserName: userName}
	UserMap[userId] = user

}
