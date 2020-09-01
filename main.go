package main

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
  "net/url"
	"time"
  "html/template"
  "strconv"
)

type attributes struct {
  Concelho string `json:"Concelho"`
  DataConc int `json:"Data_Conc"`
  DataParsed string
  RecuperadosConc int `json:"Recuperados_Conc"`
  ConfirmadosAcumuladoConc int `json:"ConfirmadosAcumulado_Conc"`
  ConfirmadosNovosConc int `json:"ConfirmadosNovos_Conc"`
}

type fields struct {
  Name string `json:"name"`
  Alias string `json:"alias"`
}

type features struct {
  Attributes attributes `json:"attributes"`
}

type concelhosData struct {
  Features []features `json:"features"`
}

type covidData2 struct {
	Number int `json:"number"`
	UpdatedParsed string
	Updated int `json:"updated"`
	Country string `json:"country"`
	Cases int `json:"cases"`
  TodayCases int `json:"todayCases"`
  Deaths int `json:"deaths"`
  TodayDeaths int `json:"todayDeaths"`
  Recovered int `json:"recovered"`
  TodayRecovered int `json:"todayRecovered"`
  Active int `json:"active"`
  Critical int `json:"critical"`
  CasesPerOneMillion float64 `json:"casesPerOneMillion"`
  DeathsPerOneMillion float64 `json:"deathsPerOneMillion"`
  Tests int `json:"tests"`
  TestsPerOneMillion float64 `json:"testsPerOneMillion"`
  Population int `json:"population"`
  Continent string `json:"continent"`
  OneCasePerPeople int `json:"oneCasePerPeople"`
  OneDeathPerPeople int `json:"oneDeathPerPeople"`
  OneTestPerPeople int `json:"oneTestPerPeople"`
  ActivePerOneMillion float64 `json:"activePerOneMillion"`
  RecoveredPerOneMillion float64 `json:"recoveredPerOneMillion"`
  CriticalPerOneMillion float64 `json:"criticalPerOneMillion"`
}

type covidData struct {
	Data string `json:"data"`
	DataDados string `json:"data_dados"`
	Confirmados float64 `json:"confirmados"`
	ConfirmadosArsnorte float64 `json:"confirmados_arsnorte"`
  ConfirmadosArscentro float64 `json:"confirmados_arscentro"`
  ConfirmadosArslvt float64 `json:"confirmados_arslvt"`
  ConfirmadosArsalentejo float64 `json:"confirmados_arsalentejo"`
  ConfirmadosArsalgarve float64 `json:"confirmados_arsalgarve"`
  ConfirmadosAcores float64 `json:"confirmados_acores"`
  ConfirmadosMadeira float64 `json:"confirmados_madeira"`
  ConfirmadosEstrangeiro float64 `json:"confirmados_estrangeiro"`
  ConfirmadosNovos float64 `json:"confirmados_novos"`
  Recuperados float64 `json:"recuperados"`
  Obitos float64 `json:"obitos"`
  Internados float64 `json:"internados_uci"`
  Lab float64 `json:"lab"`
  Suspeitos float64 `json:"suspeitos"`
  Vigilancia float64 `json:"vigilancia"`
  NConfirmados float64 `json:"n_confirmados"`
  CadeiasTransmissao float64 `json:"cadeias_transmissao"`
  TransmissaoImportada float64 `json:"transmissao_importada"`
  Confirmados09f float64 `json:"confirmados_0_9_f"`
  Confirmados09m float64 `json:"confirmados_0_9_m"`
  Confirmados1019f float64 `json:"confirmados_10_19_f"`
  Confirmados1019m float64 `json:"confirmados_10_19_m"`
  Confirmados2029f float64 `json:"confirmados_20_29_f"`
  Confirmados2029m float64 `json:"confirmados_20_29_m"`
  Confirmados3039f float64 `json:"confirmados_30_39_f"`
  Confirmados3039m float64 `json:"confirmados_30_39_m"`
  Confirmados4049f float64 `json:"confirmados_40_49_f"`
  Confirmados4049m float64 `json:"confirmados_40_49_m"`
  Confirmados5059f float64 `json:"confirmados_50_59_f"`
  Confirmados5059m float64 `json:"confirmados_50_59_m"`
  Confirmados6069f float64 `json:"confirmados_60_69_f"`
  Confirmados6069m float64 `json:"confirmados_60_69_m"`
  Confirmados7079f float64 `json:"confirmados_70_79_f"`
  Confirmados7079m float64 `json:"confirmados_70_79_m"`
  Confirmados80Plusf float64 `json:"confirmados_80_plus_f"`
  Confirmados80Plusm float64 `json:"confirmados_80_plus_m"`
  SintomasTosse float64 `json:"sintomas_tosse"`
  SintomasFebre float64 `json:"sintomas_febre"`
  SintomasDificuldadeRespiratoria float64 `json:"sintomas_dificuldade_respiratoria"`
  SintomasCefaleia float64 `json:"sintomas_cefaleia"`
  SintomasDoresMusculares float64 `json:"sintomas_dores_musculares"`
  SintomasFraquezaGeneralizada float64 `json:"sintomas_fraqueza_generalizada"`
  ConfirmadosF float64 `json:"confirmados_f"`
  ConfirmadosM float64 `json:"confirmados_m"`
  ObitosArsnorte float64 `json:"obitos_arsnorte"`
  ObitosArscentro float64 `json:"obitos_arscentro"`
  ObitosArslvt float64 `json:"obitos_arslvt"`
  ObitosArsalentejo float64 `json:"obitos_arsalentejo"`
  ObitosArsalgarve float64 `json:"obitos_arsalgarve"`
  ObitosAcores float64 `json:"obitos_acores"`
  ObitosMadeira float64 `json:"obitos_madeira"`
  ObitosEstrangeiro float64 `json:"obitos_estrangeiro"`
  RecuperadosArsnorte float64 `json:"recuperados_arsnorte"`
  RecuperadosArscentro float64 `json:"recuperados_arscentro"`
  RecuperadosArslvt float64 `json:"recuperados_arslvt"`
  RecuperadosArsalentejo float64 `json:"recuperados_arsalentejo"`
  RecuperadosArsalgarve float64 `json:"recuperados_arsalgarve"`
  RecuperadosAcores float64 `json:"recuperados_acores"`
  RecuperadosMadeira float64 `json:"recuperados_madeira"`
  RecuperadosEstrangeiro float64 `json:"recuperados_estrangeiro"`
  Obitos09f float64 `json:"obitos_0_9_f"`
  Obitos09m float64 `json:"obitos_0_9_m"`
  Obitos1019f float64 `json:"obitos_10_19_f"`
  Obitos1019m float64 `json:"obitos_10_19_m"`
  Obitos2029f float64 `json:"obitos_20_29_f"`
  Obitos2029m float64 `json:"obitos_20_29_m"`
  Obitos3039f float64 `json:"obitos_30_39_f"`
  Obitos3039m float64 `json:"obitos_30_39_m"`
  Obitos4049f float64 `json:"obitos_40_49_f"`
  Obitos4049m float64 `json:"obitos_40_49_m"`
  Obitos5059f float64 `json:"obitos_50_59_f"`
  Obitos5059m float64 `json:"obitos_50_59_m"`
  Obitos6069f float64 `json:"obitos_60_69_f"`
  Obitos6069m float64 `json:"obitos_60_69_m"`
  Obitos7079f float64 `json:"obitos_70_79_f"`
  Obitos7079m float64 `json:"obitos_70_79_m"`
  Obitos80Plusf float64 `json:"obitos_80_plus_f"`
  Obitos80Plusm float64 `json:"obitos_80_plus_m"`
  ObitosF float64 `json:"obitos_f"`
  ObitosM float64 `json:"obitos_m"`
  ConfirmadosDesconhecidosM float64 `json:"confirmados_desconhecidos_m"`
  ConfirmadosDesconhecidosF float64 `json:"confirmados_desconhecidos_f"`
  Ativos float64 `json:"ativos"`

}

type data struct {
  Covid covidData
  Concelhos concelhosData
  Covid2 covidData2
}

func main() {

	//We tell Go exactly where we can find our html file.
	//We ask Go to parse the html file (Notice the relative path).
	//We wrap it in a call to template.Must() which handles any errors
	//and halts if there are fatal errors
   
	 templates := template.Must(template.ParseFiles("templates/index.html", "templates/regional.html", "templates/concelho.html"))
	 //Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
   // a handle that looks in the static directory, go then uses the "/static/" as a url that our
   //html can refer to when looking for our css and other files. 

	http.Handle("/static/", //final url can be anything
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

    covid := covidData{}
    getCovidData(&covid)
    concelhos := concelhosData{}
    getCovidData2(&concelhos)
    covid2 := covidData2{}
    getCovidData3(&covid2)

    data := data{Covid: covid, Concelhos: concelhos, Covid2: covid2}

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
  })
  http.HandleFunc("/regional/" , func(w http.ResponseWriter, r *http.Request) {

    covid := covidData{}
    getCovidData(&covid)
    concelhos := concelhosData{}
    getCovidData2(&concelhos)

    data := data{Covid: covid, Concelhos: concelhos}

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "regional.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
  })
  http.HandleFunc("/concelho/" , func(w http.ResponseWriter, r *http.Request) {

    covid := covidData{}
    getCovidData(&covid)
    concelhos := concelhosData{}
    getCovidData2(&concelhos)

    data := data{Covid: covid, Concelhos: concelhos}

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "concelho.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

   //Start the web server, set the port to listen to 8080. Without a path it assumes localhost
   //Print any errors from starting the webserver using fmt
   fmt.Println("Listening");
	 fmt.Println(http.ListenAndServe(":9090", nil));


	//  lines, err := ReadCsv("covid19pt-data/data.csv")
  //   if err != nil {
  //       panic(err)
  //   }

  //   // Loop through lines & turn into object
  //   for _, line := range lines {
  //       data := CsvLine{
  //           Column1: line[0],
  //           Column2: line[1],
  //       }
  //       fmt.Println(data.Column1 + " " + data.Column2)
  //   }
}

func getCovidData(covid *covidData){
  // url := "https://disease.sh/v3/covid-19/countries/portugal"
  url2 := "https://covid19-api.vost.pt/Requests/get_last_update"

  spaceClient := http.Client{
    Timeout: time.Second * 10, // Timeout after 10 seconds
  }

  req, err := http.NewRequest(http.MethodGet, url2, nil)
  if err != nil {
    log.Fatal(err)
  }

  //req.Header.Set("User-Agent", "spacecount-tutorial")

  res, getErr := spaceClient.Do(req)
  if getErr != nil {
    log.Fatal(getErr)
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    log.Fatal(readErr)
  }

  jsonErr := json.Unmarshal(body, &covid)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }
}

func getCovidData2(concelhos *concelhosData){
  data := url.Values {
    "f": {"json"},
    "where": {"1=1"},
    "outFields": {"Concelho", "ConfirmadosAcumulado_Conc", "ConfirmadosNovos_Conc", "Data_Conc"},
    "returnGeometry": {"false"},
  }

  res, err := http.PostForm("https://services.arcgis.com/CCZiGSEQbAxxFVh3/ArcGIS/rest/services/COVID_Concelhos_ConcelhosDetalhes/FeatureServer/0/query", data)

  if err != nil {
      panic(err)
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    log.Fatal(readErr)
  }

  jsonErr := json.Unmarshal(body, &concelhos)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }

  for i := 0; i < len(concelhos.Features); i++ {
    concelhos.Features[i].Attributes.Concelho = strings.Title(strings.ToLower(concelhos.Features[i].Attributes.Concelho))
    concelhos.Features[i].Attributes.DataParsed = tsToDate(concelhos.Features[i].Attributes.DataConc)
  }

}

func getCovidData3(covid *covidData2){
  url := "https://disease.sh/v3/covid-19/countries/portugal"

  spaceClient := http.Client{
    Timeout: time.Second * 10, // Timeout after 10 seconds
  }

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }

  //req.Header.Set("User-Agent", "spacecount-tutorial")

  res, getErr := spaceClient.Do(req)
  if getErr != nil {
    log.Fatal(getErr)
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    log.Fatal(readErr)
  }

  jsonErr := json.Unmarshal(body, &covid)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }

  covid.UpdatedParsed = tsToDate(covid.Updated)
}


func tsToDate(ts int) string {
  s := strconv.Itoa(ts)
  if len(s) < 10 {
    return "sem informação"
  }
  s = string(s[0:10])
  if n, err := strconv.Atoi(s); err == nil {
    return time.Unix(int64(n), 0).Format("2006-01-02")
  }
  return "sem informação"
}