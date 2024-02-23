package util

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	cRand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	_ "image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/nfnt/resize"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model/errors"
)

const (
	alphabet  = "abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	minLength = 5
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenRandomInRange return a random int in a range
func GenRandomInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// GenRandomFloatInRange return a random float in a range
func GenRandomFloatInRange(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// SplitAndTrimSpaceString is a helper for split and strim space from the results
func SplitAndTrimSpaceString(s string, sep string) []string {
	if s == "" {
		return nil
	}

	if sep == "" {
		return []string{strings.TrimSpace(s)}
	}

	l := strings.Split(s, sep)
	rs := []string{}
	for i := range l {
		tmp := strings.TrimSpace(l[i])
		if tmp != "" {
			rs = append(rs, tmp)
		}
	}
	return rs
}

// RandomString generate random string with lenght
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ValidateEmail validate a string is email by regular expression
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

// ValidatePhone validate a string is phone by regular expression
func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return re.MatchString(phone)
}

// GenerateSaltedPassword generate salted password from string and salt
// Return salted string
func GenerateSaltedPassword(password, salt string, loops int) (string, error) {
	salted := salt
	passwd := password

	r := regexp.MustCompile(`^\$([0-9]+)\$(.*)`)
	subStrs := r.FindStringSubmatch(salted)
	if len(subStrs) == 3 {
		i, err := strconv.Atoi(subStrs[1])
		if err != nil {
			return "", err
		}
		loops = i
		salted = subStrs[2]
	}

	for i := 0; i < loops; i++ {
		h := sha1.New()
		h.Write([]byte(salted + passwd))
		passwd = fmt.Sprintf("%x", h.Sum(nil))
	}

	return fmt.Sprintf("$%d$%s", loops, salted+passwd), nil
}

// HashNumber hash a number to string by sha1 algolithm
// Return hashed string
func HashNumber(val int64) string {
	b := []byte(strconv.FormatInt(val, 10))
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// CopyMap make a copy from map
// Return map struct data
func CopyMap(src map[string]interface{}) map[string]interface{} {
	rs := map[string]interface{}{}
	for k, v := range src {
		rs[k] = v
	}

	return rs
}

// ParseErrorCode parse error code from errors.Error
func ParseErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch arg := err.(type) {
	case *errors.Error:
		return int(arg.Code)

	case error:
		return http.StatusInternalServerError

	default:
		return http.StatusOK
	}
}

// HandleError handler error from errors.Error
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch arg := err.(type) {

	case *errors.Error:
		c.JSON(int(arg.Code), arg)

	case errors.Error:
		c.JSON(int(arg.Code), arg)

	case error:
		c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    http.StatusInternalServerError,
			Message: arg.Error(),
		})
	}

}

// GenUniqueCode return a unique string
func GenUniqueCode() string {
	rs, err := gonanoid.Generate(alphabet, minLength)
	if err != nil {
		return RandomString(minLength)
	}
	return rs

}

// HashBase64String return a hash string
func HashBase64String(val string) string {
	h := sha256.New()
	h.Write([]byte(val))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// TimePart get froms and tos from time duration
func TimePart(from, to time.Time, durationType string) ([]time.Time, []time.Time) {
	var froms, tos []time.Time
	i := 0
	for {
		tempFrom := time.Date(from.Year(), from.Month(), from.Day()+i, 0, 0, 0, 0, from.Location())
		tempTo := time.Date(from.Year(), from.Month(), from.Day()+1+i, 0, 0, 0, 0, from.Location())
		if durationType == "month" {
			tempFrom = time.Date(from.Year(), from.Month()+time.Month(i), 1, 0, 0, 0, 0, from.Location())
			tempTo = time.Date(from.Year(), from.Month()+1+time.Month(i), 1, 0, 0, 0, 0, from.Location())
		}
		froms = append(froms, tempFrom)
		tos = append(tos, tempTo)
		i++
		if tempTo.Equal(to) || (tempFrom.Month() == to.Month() && tempFrom.Year() == to.Year()) {
			if durationType == "day" {
				if tempTo.Equal(to) || tempFrom.Day() == to.Day() {
					break
				}
				continue
			}
			break
		}
	}
	return froms, tos
}

func Encrypt(text, key string) (string, error) {
	tB := []byte(text)
	kB := []byte(key)
	c, err := aes.NewCipher(kB)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cRand.Reader, nonce); err != nil {
		return "", err
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, tB, nil)), nil
}

func Decrypt(text, key string) (string, error) {
	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	kB := []byte(key)
	c, err := aes.NewCipher(kB)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.NewStringError("invalid cipher", http.StatusBadRequest)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func CheckAndResizeImg(fileName string) (string, int, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return fileName, 0, 0, err
	}

	img, _, err := image.Decode(file)

	height := img.Bounds().Dx()
	width := img.Bounds().Dy()
	if err != nil {
		return fileName, 0, 0, err
	}
	file.Close()

	if height >= 2048 || width >= 2048 {
		// resize to width 1000 using Lanczos resampling
		m := resize.Resize(2048, 2048, img, resize.Lanczos3)
		resizeFileName := strings.Replace(fileName, ".", "_resize.", -1)
		resizeFile, err := os.Create(resizeFileName)
		if err != nil {
			return fileName, 0, 0, err
		}

		err = jpeg.Encode(resizeFile, m, nil)
		if err != nil {
			return fileName, 0, 0, err
		}

		fileName = resizeFileName
		defer resizeFile.Close()
	}

	return fileName, height, width, nil
}

func ChangeFormatIpfs(ipfsUrl string) string {
	if strings.Contains(ipfsUrl, "cloudflare-ipfs.com") {
		cloufareConponent := strings.Split(ipfsUrl, "cloudflare-ipfs.com")
		ipfsUrl = cloufareConponent[0] + "ipfs.io" + cloufareConponent[1]
	}

	if strings.Contains(ipfsUrl, "ipfs.infura.io") {
		infuraComponent := strings.Split(ipfsUrl, "ipfs.infura.io")
		ipfsUrl = infuraComponent[0] + "ipfs.io" + infuraComponent[1]
	}
	return ipfsUrl
}

func GetRemainingDaysFromTS(ts int) int {
	now := time.Now()
	diff := time.Unix(int64(ts), 0).Sub(now)
	return int(diff.Hours() / 24)
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func DownloadFile(URL, imageFile string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Create a empty file
	file, err := os.Create(imageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func MaskAddress(str string) string {
	if len(str) > 7 {
		return fmt.Sprintf("%s***%s", str[:7], str[len(str)-7:])
	}

	return str
}

var discordEmojiCache = make(map[string]string)

func GetDiscordEmoji(session *discordgo.Session, guildID string, emojiName string) (string, error) {

	emoji, ok := discordEmojiCache[emojiName]
	if ok {
		return emoji, nil
	}

	emojis, err := session.GuildEmojis(guildID)
	if err != nil {
		return "", err
	}

	emojiMsg := ""
	for _, e := range emojis {
		if emojiName == e.Name {
			emojiMsg = e.MessageFormat()
		}
		discordEmojiCache[e.Name] = e.MessageFormat()
	}

	if emojiMsg == "" {
		return "", fmt.Errorf("emoji %s not found", emojiName)
	}

	return emojiMsg, nil
}

func FormatDiffTimeToHumanReadable(a, b time.Time) (result string) {
	var year, month, day, hour, min, sec int
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	if year > 0 {
		result = result + fmt.Sprintf("%d year ", year)
	}
	if month > 0 {
		result = result + fmt.Sprintf("%d month ", month)
	}
	if day > 0 {
		result = result + fmt.Sprintf("%d day ", day)
	}
	if hour > 0 {
		result = result + fmt.Sprintf("%d hour ", hour)
	}
	// min, sec if needed

	return
}

func WeiToEther(wei *big.Int, decimals ...int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	var e *big.Float
	if len(decimals) == 0 {
		e = big.NewFloat(params.Ether)
	} else {
		e = big.NewFloat(math.Pow(10, float64(decimals[0])))
	}
	return f.Quo(fWei.SetInt(wei), e)
}

func StringWeiToEther(stringWei string, decimals int) *big.Float {
	if decimals == 0 {
		decimals = 18
	}
	wei := new(big.Int)
	wei.SetString(stringWei, 10)
	return WeiToEther(wei, decimals)
}

func TrimAddressFromLog(s string) string {
	return strings.ReplaceAll(s, "0x000000000000000000000000", "0x")
}

type getSenderByTxHashResp struct {
	From *common.Address `json:"from,omitempty"`
}

func GetSenderByTxHash(rpcEndpoint string, txHash common.Hash) (common.Address, error) {
	rpc, err := rpc.DialContext(context.Background(), rpcEndpoint)
	if err != nil {
		return common.Address{}, err
	}

	var json getSenderByTxHashResp
	if err = rpc.CallContext(context.Background(), &json, "eth_getTransactionByHash", txHash); err != nil {
		return common.Address{}, err
	}

	if json.From == nil {
		return common.Address{}, fmt.Errorf("no sender found")
	}

	return *json.From, nil
}

func Uint8ToIntPointer(u uint8) *int {
	i := int(u)
	return &i
}

func FetchData(url string, parseForm interface{}) (int, error) {
	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Get(url)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			for i := 0; i < 2; i++ {
				logger.NewLogrusLogger().Fields(logger.Fields{"url": url}).Infof(fmt.Sprintf("context deadline exceeded for service, retrying %dth ...", i+1))
				resp, err = client.Get(url)
				if err == nil {
					break
				}
			}
		} else {
			return http.StatusInternalServerError, err
		}
	}

	if resp == nil {
		return http.StatusInternalServerError, fmt.Errorf("cannot get data from url: %s", url)
	}

	statusCode := resp.StatusCode
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return statusCode, err
	}

	resp.Body.Close()

	return statusCode, json.Unmarshal(b, parseForm)
}

func GetMaxFloat64(arr []float64) float64 {
	max := arr[0]
	for _, ele := range arr {
		if ele > max {
			max = ele
		}
	}

	return RoundFloat(max, 2)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetNullUUID(id string) uuid.NullUUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Error("uuid invalid")
	}
	if id == "" {
		return uuid.NullUUID{Valid: false}
	}
	nullid := uuid.NullUUID{UUID: uid, Valid: true}
	return nullid
}

func ConvertToFloat(amount string, decimal int) float64 {
	tmp, _ := strconv.ParseInt(amount, 10, 64)

	dec := float64(decimal)
	amnt := float64(tmp)
	value := amnt * math.Pow(10, -dec)
	return value
}

func SecondsToDays(sec int) int {
	return sec / 86400
}

func GetStringBetweenParentheses(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			return s[i+1 : j]
		}
	}
	return s
}

func SetRequestBody(c *gin.Context, structBody interface{}) {
	ctx := *c
	json, err := json.Marshal(structBody)
	if err != nil {
		log.Error("cannot encode body")
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(json))
	// 1. set new header
	ctx.Request.Header.Set("Content-Length", strconv.Itoa(len(json)))
	// 2. also update this field
	ctx.Request.ContentLength = int64(len(json))
	c = &ctx
}

func MinuteLeftUntil(startTime, endTime time.Time) float64 {
	var minutes float64 = 0
	if startTime.Before(endTime) {
		minutes = endTime.Sub(startTime).Minutes()
	}
	return minutes
}

func MinInt(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func Shuffle[T any](list []T) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func NumberPostfix(num int) string {
	var postfix string
	switch num % 10 {
	case 1:
		postfix = "st"
	case 2:
		postfix = "nd"
	case 3:
		postfix = "rd"
	default:
		postfix = "th"
	}
	return postfix
}

func RemoveAt[T any](list []T, idx int) []T {
	return append(list[:idx], list[idx+1:]...)
}

type SendRequestQuery struct {
	Method      string // default = GET
	URL         string
	Response    interface{}
	ErrResponse interface{}
	Headers     map[string]string
	Body        io.Reader
}

func SendRequest(q SendRequestQuery) (int, error) {
	if q.Method == "" {
		q.Method = http.MethodGet
	}
	client := &http.Client{Timeout: time.Second * 30}
	req, _ := http.NewRequest(q.Method, q.URL, q.Body)
	for k, v := range q.Headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			for i := 0; i < 3; i++ {
				logger.NewLogrusLogger().Fields(logger.Fields{"q": q}).Infof(fmt.Sprintf("context deadline exceeded for service, retrying %dth ...", i+1))
				res, err = client.Do(req)
				if err == nil {
					break
				}
			}
		} else {
			return http.StatusInternalServerError, err
		}
	}

	if res == nil {
		return http.StatusInternalServerError, fmt.Errorf("cannot get data from url: %s", q.URL)
	}

	statusCode := res.StatusCode

	// Handle 503 error first before unmarshal response, because 503 error response sometimes is not json -> cannot unmarshal -> let unmarshal error is returned instead of 503 error
	if statusCode == http.StatusServiceUnavailable {
		return statusCode, nil
	}

	// Handle error response when it is an HTML page, can consider to parse error message from HTML page later
	contentType := res.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") {
		// Handle HTML error response
		htmlBody, err := io.ReadAll(res.Body)
		if err != nil {
			return statusCode, err
		}

		// Just logging the error response and return response as empty -> get no data
		logger.NewLogrusLogger().Fields(logger.Fields{"q": q, "htmlBody": string(htmlBody)}).Info("Error response from service is not json")
		return statusCode, nil
	}

	if q.Response != nil {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return statusCode, err
		}
		if err := json.Unmarshal(bytes, q.Response); err != nil {
			return statusCode, err
		}
		return statusCode, nil
	}

	res.Body.Close()

	return statusCode, nil
}

func ParseSnapshotURL(url string) string {
	//https://snapshot.org/#/bitdao.eth
	if strings.Contains(url, "snapshot.org/#/") {
		args := strings.Split(url, "/")
		return args[len(args)-1]
	}
	return url
}

func FloatToBigInt(val float64, decimals int64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	decimalValue := math.Pow(10, float64(decimals))

	coin := new(big.Float)
	coin.SetInt(big.NewInt(int64(math.Floor(decimalValue))))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func Capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// TODO(trkhoi): temp keep this until i refactor vault noti -> mochi-notification
func TokenEmoji(token string) string {
	tokenEmojis := map[string]string{
		"FTM":    "<:ftm:967285237686108212>",
		"SPIRIT": "<:spirit:967285237962924163>",
		"TOMB":   "<:tomb:967285237904179211>",
		"REAPER": "<:reaper:967285238306857063>",
		"BOO":    "<:boo:967285238042599434>",
		"SPELL":  "<:spell:967285238063587358>",
		"BTC":    "<:btc:967285237879013388>",
		"ETH":    "<:eth:991657409082830858>",
		"WETH":   "<:weth:1052849700019109889>",
		"BNB":    "<:bnb:972205674715054090>",
		"CAKE":   "<:cake:972205674371117126>",
		"OP":     "<:op:1002151912403107930>",
		"USDT":   "<:usdt:1005010747308396544>",
		"USDC":   "<:usdc:1005010675342520382>",
		"ADA":    "<:ada:1005010608443359272>",
		"XRP":    "<:xrp:1005010559856554086>",
		"BUSD":   "<:busd:1005010097535197264>",
		"DOT":    "<:dot:1005009972716908554>",
		"DOGE":   "<:doge:1004962950756454441>",
		"DAI":    "<:dai:1005009904433647646>",
		"MATIC":  "<:matic:1037985931816349746>",
		"AVAX":   "<:avax:1005009817523474492>",
		"UNI":    "<:uni:1005012087967334443>",
		"SHIB":   "<:shib:1005009723277463703>",
		"TRX":    "<:trx:1005009394209128560>",
		"WBTC":   "<:wbtc:1005009348956790864>",
		"ETC":    "<:etc:1005009314802569277>",
		"LEO":    "<:leo:1005009244187263047>",
		"LTC":    "<:ltc:1005009185940963380>",
		"FTT":    "<:ftt:1005009144044064779>",
		"CRO":    "<:cro:1005009127937949797>",
		"LINK":   "<:link:1005008904205385759>",
		"NEAR":   "<:near:1005008870038589460>",
		"ATOM":   "<:atom:1005008855111049216>",
		"XLM":    "<:xlm:1005008839139151913>",
		"XMR":    "<:xmr:1005008819866312724>",
		"BCH":    "<:bch:1005008800106942525>",
		"APE":    "<:ape:1005008782486675536>",
		"DFG":    "<:dfg:1007157463256145970>",
		"ICY":    ":ice_cube:",
		"CARROT": ":carrot:",
		"BUTT":   "<:butt:1007247521468403744>",
		"WDOGE":  "<:wdoge:1010512669448605756>",
		"REN":    "<:ren:1037985602202779690>",
		"MANA":   "<:mana:1037985604010508360>",
		"COMP":   "<:comp:1037985570724528178>",
		"YFI":    "<:yfi:1037985592971116564>",
		"BAT":    "<:bat:1037985578341371964>",
		"AAVE":   "<:aave:1037985567146774538>",
		"BNT":    "<:bnt:1037985589355626517>",
		"MKR":    "<:mkr:1037985596964081696>",
		"ANC":    "<:anc:1037985575334051901>",
		"BRUSH":  "<:brush:1037985582162378783>",
		"SOL":    "<:sol:1006838270862295080>",
		"APT":    "<:apt:1047707078183096320>",
		"RONIN":  "<:ronin:1047707529884483614>",
		"ARB":    "<:arb:1056772215477112862>",
		"OKC":    "<:okc:1006838263165767681>",
		"ONUS":   "<:onus:1077203550075093053>",
		"SUI":    "<:sui:1077132420500951081>",
		"FBOMB":  "<:fbomb:1079669535117938788>",
		"MCLB":   "<:mlcb:1079669537408036955>",
		"BSC":    "<:bsc:972205674715054090>",
		"POL":    "<:pol:1037985931816349746>",
	}
	val, ok := tokenEmojis[token]
	// If the key exists
	if ok {
		return val
	}
	return "<:ftm:967285237686108212>"

}

func CheckKeyInMap(key string, m interface{}) bool {
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			if k.String() == key {
				return true
			}
		}
	}
	return false
}

func ValidateNumberSeries(s string) bool {
	if s == "" {
		return false
	}

	matched, err := regexp.MatchString("^[0-9]*$", s)
	if err != nil {
		return false
	}

	return matched
}

func ValidateFileMarkdown(s string) bool {
	if s == "" {
		return false
	}

	matched, err := regexp.MatchString(".*?.md", s)
	if err != nil {
		return false
	}

	return matched
}
