package main

import(
    "math"
    "time"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/load"
    "encoding/json"
    "net/http"
    "log"
    "html/template"
    "os"
)

type response struct {
    Cpu  float64  `json:"cpu"`
    Load  float64    `json:"load"`
    Mem float64 `json:"mem"`
    Version string `json:"version"`
}

type Page struct {
    Title string
    Body  []byte
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func statHandler(rw http.ResponseWriter, r *http.Request) {
    version := "0.1"
    v, _ := mem.VirtualMemory()
    c, _ := cpu.Percent((500 * time.Millisecond), false)
    l, _ := load.Avg()

    // format everything using the struct
    // res1 := &response{
    //     Cpu: c[0],
    //     Load: l.Load1,
    //     Mem: v.UsedPercent}
    // fmt.Println(res1)
    //log.Println(res1.cpu)
    log.Println("returning JSON")
    json.NewEncoder(rw).Encode(response{Cpu: toFixed(c[0], 2),
                                        Load: toFixed(l.Load1, 2),
                                        Mem: toFixed(v.UsedPercent, 2),
                                        Version: version})
}

func dashHandler(w http.ResponseWriter, r *http.Request) {
    title, err := os.Hostname()
    if err != nil {
		panic(err)
	}
    p := &Page{Title: title}
    t, _ := template.ParseFiles("www/dashboard.html")
    log.Println("Attempting to render dashboard")
    t.Execute(w, p)
}

func main() {
    address := ":1234"
    r := http.NewServeMux()
    r.HandleFunc("/__healthcheck__", statHandler)
    r.HandleFunc("/", dashHandler)
    log.Println("Healthcheck Service request at: " + address)
    log.Println("Please visit http://localhost:1234/")
    //log.Println("hostname: " + os.Hostname())
    log.Println(http.ListenAndServe(address, r))

    // mem - almost every return value is a struct
    // fmt.Printf("Mem: %f\n", v.UsedPercent)

    // cpu  
    // fmt.Printf("CPU: %f\n", c[0])

    // load
    // fmt.Printf("Load: %f\n", l.Load1)

}
