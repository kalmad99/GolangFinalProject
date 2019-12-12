package main

import (
	"database/sql"
	"fmt"
	"net/smtp"

	//"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	//_ "github.com/lib/pq"
	_ "./menu/github.com/go-sql-driver/mysql"

	"./entity"
	"./menu/repository"
	"./menu/service"
)


var (
	name string
	email string
	pass string
)

type FirstPageInfo struct {
	First string
	Email string
	Password string
}

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*.html"))
var productService *service.ProductService

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	products, err := productService.Products()
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.layout", products)
	if err!=nil{
		panic(err)
	}
}
func indexMob(w http.ResponseWriter, r *http.Request) {

	mobile, err := productService.MobProducts()
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.layout", mobile)
	if err!=nil{
		panic(err)
	}
}
func indexComp(w http.ResponseWriter, r *http.Request) {

	computers, err := productService.CompProducts()
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.layout", computers)
	if err!=nil{
		panic(err)
	}
}
func indexLap(w http.ResponseWriter, r *http.Request) {

	laptops, err := productService.LapProducts()
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.layout", laptops)
	if err!=nil{
		panic(err)
	}
}
func indexCam(w http.ResponseWriter, r *http.Request) {

	cameras, err := productService.CamProducts()
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.layout", cameras)
	if err!=nil{
		panic(err)
	}
}

func searchProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		res := r.URL.Query().Get("search")

		if len(res)==0{
			http.Redirect(w, r, "/", 303)
		}
		results, err := productService.SearchProduct(res)

		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(w, "searchresults.layout", results)

		if err != nil{
			panic(err.Error())
		}

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func productDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		pro, err := productService.Product(id)

		if err != nil {
			panic(err)
		}

		_ = tmpl.ExecuteTemplate(w, "productdetail.layout", pro)
	}
}

//func about(w http.ResponseWriter, r *http.Request) {
//	tmpl.ExecuteTemplate(w, "about.layout", nil)
//}
//
//func menu(w http.ResponseWriter, r *http.Request) {
//	tmpl.ExecuteTemplate(w, "menu.layout", nil)
//}
//
//func contact(w http.ResponseWriter, r *http.Request) {
//	tmpl.ExecuteTemplate(w, "contact.layout", nil)
//}
//
func seller(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "seller.index.layout", nil)
}

func indexProducts(w http.ResponseWriter, r *http.Request) {
	products, err := productService.Products()
	if err != nil {
		panic(err)
	}
	_ = tmpl.ExecuteTemplate(w, "seller.products.layout", products)
}

func sellerNewProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.Product{}
		ctg.Name = r.FormValue("name")
		ctg.Quantity, _ = strconv.Atoi(r.FormValue("quantity"))
		ctg.ItemType = r.FormValue("type")
		ctg.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
		ctg.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		err = productService.StoreProduct(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/seller/products", http.StatusSeeOther)

	} else {

		err := tmpl.ExecuteTemplate(w, "seller.product.new.layout", nil)

		if err!=nil{
			panic(err)
		}

	}
}

func ratingPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		idRaw := req.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		pro, err := productService.Product(id)

		if err != nil {
			panic(err)
		}
		_ = tmpl.ExecuteTemplate(w, "ratings.html", pro)
	} else if req.Method == http.MethodPost {

		prod := entity.Product{}
		prod.ID, _ = strconv.Atoi(req.FormValue("id"))
		prod.Rating, _ = strconv.ParseFloat(req.FormValue("rate"), 64)

		_, err := productService.RateProduct(prod)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, req, "/", http.StatusSeeOther)

	}else {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

//func ratingafter(w http.ResponseWriter, req *http.Request){
//	if req.Method == http.MethodGet {
//		idRaw := req.URL.Query().Get("id")
//		id, err := strconv.Atoi(idRaw)
//
//		//rateRaw := req.URL.Query().Get("rating")
//		//rate, err := strconv.ParseFloat(rateRaw, 64)
//		if err != nil {
//			panic(err)
//		}
//		pro, err := productService.Product(id)
//
//		prowithrate, err := productService.RateProduct(pro)
//
//		if err != nil {
//			panic(err)
//		}
//
//		err = tmpl.ExecuteTemplate(w, "productdetail.layout", prowithrate)
//		if err!=nil{
//			panic(err.Error())
//		}
//
//	}
//	//if req.Method == http.MethodPost{
//	//	prod := entity.Product{}
//	//	prod.Rating, _ = strconv.ParseFloat(req.FormValue("rate"), 64)
//	//	prod.ID, _ = strconv.Atoi(req.FormValue("id"))
//	//
//	//
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//if req.Method == http.MethodPost{
//	//	Rating := req.FormValue("rate")
//	//
//	//	err := tmpl.ExecuteTemplate(w, "ratingafter.html", Rating)
//	//	if err!=nil{
//	//		panic(err.Error())
//	//	}
//	//}
//}
func sellerUpdateProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := productService.Product(id)

		if err != nil {
			panic(err)
		}

		_ = tmpl.ExecuteTemplate(w, "seller.products.update.layout", cat)

	} else if r.Method == http.MethodPost {

		prod := entity.Product{}
		prod.ID, _ = strconv.Atoi(r.FormValue("id"))
		prod.Name = r.FormValue("name")
		prod.Description = r.FormValue("description")
		prod.Image = r.FormValue("image")

		mf, _, err := r.FormFile("catimg")

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		writeFile(&mf, prod.Image)

		err = productService.UpdateProduct(prod)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/seller/products", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/seller/products", http.StatusSeeOther)
	}

}

func sellerDeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = productService.DeleteProduct(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/seller/products", http.StatusSeeOther)
}

func regist(w http.ResponseWriter, req *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "Registrationform.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "login.html", nil)
}

func regist2part1 (w http.ResponseWriter, req *http.Request){
	if req.Method != "POST"{
		http.Redirect(w, req, "/registration", http.StatusSeeOther)
		return
	}

	name = req.FormValue("name")
	email = req.FormValue("email")
	pass = req.FormValue("pass")

	info := FirstPageInfo {
		First:name,
		Email:email,
		Password:pass,
	}

	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "kalemesfin12go@gmail.com"
	password := "qnzfgwbnaxykglvu"
	emailReceiver := email

	emailAuth := smtp.PlainAuth(
		"",
		emailSender,
		password,
		hostURL,
	)

	msg := []byte("To: " + emailReceiver + "\r\n" +
		"Subject: " + "Hello " + name + "\r\n" +
		"This is your OTP. 123456789")

	err:=	smtp.SendMail(
		hostURL + ":" + hostPort,
		emailAuth,
		emailSender,
		[]string{emailReceiver},
		msg,
	)

	if err != nil{
		fmt.Print("Error: ", err)
	}
	fmt.Print("Email Sent")

	_ = tmpl.ExecuteTemplate(w, "Registrationformpart2.html", info)
}

func regist2part2(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST"{
		http.Redirect(w, req, "/registrationpart2", http.StatusSeeOther)
		return
	}

	otp := req.FormValue("otpfield")

	info2 := FirstPageInfo{
		First: name,
		Email: email,
		Password: pass,
	}

	if otp == "123456789" {
		_ = tmpl.ExecuteTemplate(w, "index.html", info2)
	} else{
		fmt.Print("Wrong otp")
		http.Redirect(w, req, "/registrationpart2", http.StatusSeeOther)
	}
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "delivery", "web", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

func main() {

	dbDriver := "mysql"
	dbName := "golangtrialdb2"
	dbUser := "root"
	dbPass := ""
	dbconn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	//fmt.Println("Successfully connected to Mysql")
	//return dbconn

	//dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$w0rdD2@localhost/golangtrialdb?sslmode=disable")
	//
	//if err != nil {
	//	panic(err)
	//}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	proRepo := repository.NewPsqlProductRepository(dbconn)
	productService = service.NewProductService(proRepo)

	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/mobile", indexMob)
	http.HandleFunc("/laptop", indexLap)
	http.HandleFunc("/computer", indexComp)
	http.HandleFunc("/camera", indexCam)
	http.HandleFunc("/searchProducts", searchProducts)
	http.HandleFunc("/detail", productDetail)
	http.HandleFunc("/seller", seller)
	http.HandleFunc("/seller/products", indexProducts)
	http.HandleFunc("/seller/products/new", sellerNewProducts)
	http.HandleFunc("/seller/products/update", sellerUpdateProducts)
	http.HandleFunc("/seller/products/delete", sellerDeleteProduct)
	http.HandleFunc("/rate", ratingPage)
	http.HandleFunc("/registrationpage", regist)
	http.HandleFunc("/registrationprocess1", regist2part1)
	http.HandleFunc("/registrationprocess2", regist2part2)
	http.HandleFunc("/login", login)
	_ = http.ListenAndServe(":8181", nil)

}
