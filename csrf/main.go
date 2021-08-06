package main

import (
	"fmt"

	"github.com/justinas/nosurf"
)

func main() {
	fmt.Println(nosurf.VerifyToken("AHymZEQoX396BOMNGYjkgsSqmCNDbMTT68gEF1Mq+Sg=", "+ofd6iSeqaaCizSY33Wp0UBH7YmKtzExrqMn5+9UsHimsMYuF4cqTis+PyZW85jh1LeG2wHKKcmoSzg++T3yIQ=="))
	fmt.Println(nosurf.VerifyToken("ooFebvuIG0anMMCr5lrxjCSNHs/ZfkJcsZvSvVOGBSg=", "qOOleHbbCHGydAi+L+lBFDaUHotzlrRyv/2DZjhzuOMKYvsWjVMTNxVEyBXJs7CYEhkARKro9i4OZlHba/W9yw=="))
}
