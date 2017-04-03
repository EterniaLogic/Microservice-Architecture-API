package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"time"
	"encoding/hex"
	"math/rand"
	"strings"
)

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

// Generate an auth level
func GenAuthLevel(authlevel string, id string) string{
    auth := doMd5("xLp4iTui7vsNuBMusM"+authlevel+"ar0fjbk");
    auth = doSha256("f9m8mCiJi7Z1cGWQ2A"+auth+id+"azJ6BIrMYpqQMUrLjq");
    
    return strings.ToUpper(hex.EncodeToString([]byte(auth))); // saturate
}

func PassSum(pass string) string{
    // by hashing multiple times, the saturation of data removes the actual "value" of the password
    // this also prevents rainbow tables from getting the correct password
    sh1 := doSha1("vgqEbKvthKRVuE");
    mdx := doMd5("gwzbpY9LEaxhM8D"+pass);
    sh1 = doSha1(sh1+mdx+"VbKKAFA");
    
    return strings.ToUpper(hex.EncodeToString([]byte(doSha256("ae"+sh1+"7z"))));
}

func PreSum(pass string) string{
    // by hashing multiple times, the saturation of data removes the actual "value" of the password
    // this also prevents rainbow tables from getting the correct password
    sh1 := doSha1("PJfdhpsR3hnhKBzXxY");
    mdx := doMd5("OwZ6wn6yMS6j"+pass);
    sh1 = doSha1(sh1+mdx+"CNGA7u");
    
    return strings.ToUpper(hex.EncodeToString([]byte(doSha256("cjSvDo"+sh1+"cHY198"))));
}

func doMd5(data string) (ret string){
    sum := md5.Sum([]byte(data));
    return string(sum[:]);
}

func doSha1(data string) (ret string){
    sum := sha1.Sum([]byte(data));
    return string(sum[:]);
}

func doSha256(data string) (ret string){
    sum := sha256.Sum224([]byte(data));
    return string(sum[:]);
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func GenRandString() string{
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789");
    
	rand.Seed((time.Now().UnixNano()*11)/10);
	
    b := make([]rune, rand.Intn(40)+10)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return strings.ToUpper(hex.EncodeToString([]byte(doSha256(string(b)+"FYVfK8S7IEyOBV0"))));
}

func DateNowSQL() string{
    return time.Now().Format("2006-01-02 15:04:05");
}

// Quantify the user level
// if it is not correct, auto-patch it to "User" on the database
// returns:
//      authlevel - textual level ie: "Administrator", "User"
//      int - quantified user level
func ConvertStringToLevel(authlevel string) int{
    switch (authlevel){
        case "Administrator":
            return 120; // number both n%2=0 and n%3=0
        case "VideoModerator":
            return 93;  // number only n%3 == 0
        case "CommentModerator":
            return 80;   // number only n%2 == 0
        case "User": // Normal user
            return 5;	 // prime number
        case "Unverified": // User has not responded to registration email
            return 1;	// prime number
        case "Banned":
            return 0;
    }
    return 0;
}


func MinMaxMixer(min int, max int)(int, int){
	// reversed info?
	if(max < min){
		// do swap, min/max reversed
		tmp := min;
		min = max;
		max = tmp;
	}
	
	// Maximum # of calls
	if(max-min > 100){
		max = min+100;
	}
	
	return min,max;
}