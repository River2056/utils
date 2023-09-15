package fetch

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"utils/common"
)

func FetchJavaFiles() {
	fileList := []string{
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/DepositBonusAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/FTDBCashbackAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/LossRecoveryAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/LuckyDrawAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/NoDepositBonusAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/OptInPromoAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/ParticipantRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/Rebate100AdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/RedepositBonusAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/RedepositBonusRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/ReferenceRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/RoughExchangeRateRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/TraderCompetitionDataRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/TradeRedeemedRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/TradingBonusAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/TradingCompetitionDataRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/TradingCompetitionParticipantRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/USDTDepositAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/USDTDepositCashbackAdminRepository.java",
		"/root/excalibur/public-service/public-provider/src/main/java/com/ty/campaign/repository/WelcomeBonusAdminRepository.java",
	}

	dest := "/root/java_projects/demo/app/src/main/java/demo/app/repository/"
	if _, err := os.Stat(dest); os.IsExist(err) {
		os.MkdirAll(dest, os.ModePerm)
	}

	for _, file := range fileList {
		bytes, err := ioutil.ReadFile(file)
		common.CheckError(err)
		filename := file[strings.LastIndex(file, "/")+1:]

		newFile := fmt.Sprintf("%v%v", dest, filename)

		err = ioutil.WriteFile(newFile, bytes, 0755)
		common.CheckError(err)
	}
}
