/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Thursday 26 September 2019 - 10:10:17
** @Filename:				random.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Thursday 26 September 2019 - 10:12:07
*******************************************************************************/

package		random

import		"math/rand"
import		"time"
import		"strconv"

/*StringBytes *****************************************************************
*	Create a random string of length n
******************************************************************************/
func	StringBytes(n int) string {
    rand.Seed(time.Now().UnixNano())
	pool := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)

    for i := 0; i < n; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
	}
    return string(bytes)
}

/******************************************************************************
**	RandStringInt
**	Create a random string of length n (only numbers)
******************************************************************************/
func	StringInt(n int) string {
    rand.Seed(time.Now().UnixNano())
	pool	:=	"0123456789"
	bytes	:= make([]byte, n)

    for i	:= 0; i < n; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
	}
    return string(bytes)
}

/*Number **********************************************************************
*	Create a random int of length n
******************************************************************************/
func	RandInt(n int) int {
    rand.Seed(time.Now().UnixNano())
	pool := "0123456789"
	bytes := make([]byte, n)

    for i := 0; i < n; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
	}
	value, _ := strconv.Atoi(string(bytes))
    return value
}

/*TokenBytes ******************************************************************
*	Get a random token
******************************************************************************/
func	TokenBytes(n int) string {
    rand.Seed(time.Now().UnixNano())
	pool := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789abcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, n)

    for i := 0; i < n; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
	}
    return string(bytes)
}