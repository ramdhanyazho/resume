package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

// Data models
type Contact struct {
    Address string
    Phone   string
    Email   string
    LinkedIn string
}

type Experience struct {
    Title string
    Company string
    Location string
    Period string
    Bullets []string
}

type Certification struct {
    Name string
    Issuer string
    Year  string
    Note  string
}

type Education struct {
    Degree string
    School string
    Period string
    GPA    string
}

type PageData struct {
    Name        string
    Title       string
    Summary     string
    Skills      []string
    Contact     Contact
    Experience  []Experience
    Certifications []Certification
    Education   []Education
    Year        int
}

func main() {
    r := mux.NewRouter()

    // Static files (for favicon or custom CSS/js if needed)
    fileServer := http.FileServer(http.Dir("./static"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

    r.HandleFunc("/", HomeHandler).Methods("GET")

    // Health for Kubernetes
    r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    // Bind to PORT env (Cloud/K8s friendly)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    srv := &http.Server{
        Handler:      r,
        Addr:         ":" + port,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Printf("Starting portfolio server on :%s ...", port)
    log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // --- Fill from CV ---
    data := PageData{
        Name:  "Muhamad Ramdhany",
        Title: "Server, Network & Infrastructure Officer | IT Datacenter Specialist",
        Summary: "IT Datacenter Specialist with 3+ years of experience. Strong in Linux/Windows administration, virtualization (VMware, Nutanix), cloud (AWS/GCP/Alibaba), mail servers, and security. Proven ability to run production infrastructure and lead incident response.",
        Skills: []string{
            "Linux (all distros)", "Windows Server", "VMware vSphere", "Nutanix AHV",
            "Synology & Dell EMC Storage", "AWS (EC2/S3/Route53/VPC)", "GCP, Alibaba Cloud",
            "Active Directory / Azure Entra ID", "Office 365 Admin", "Zimbra, Exchange",
            "Networking & DNS", "Containers (Docker)", "Backup: Veeam/Acronis",
            "Monitoring & Reporting",
        },
        Contact: Contact{
            Address:  "Jl. Laksana B-II RT.004/RW.006 No.111A, Kel. Kartini, Kec. Sawah Besar, Jakarta Pusat 10750",
            Phone:    "+62 857 1443 2441",
            Email:    "ramdhanielevent@gmail.com",
            LinkedIn: "https://id.linkedin.com/in/muhamad-ramdhany-006887141",
        },
        Experience: []Experience{
            {
                Title: "Server, Network & Infrastructure Officer",
                Company: "PT. Mitsui Leasing Capital Indonesia",
                Location: "Jakarta",
                Period: "2024–Now",
                Bullets: []string{
                    "Manage VMware vCenter environments and core network performance across head office and branches.",
                    "Administer VOIP systems and coordinate operations across primary and DR datacenters.",
                },
            },
            {
                Title: "System Administrator",
                Company: "PT. Tunas Ridean",
                Location: "Jakarta",
                Period: "2022–2024",
                Bullets: []string{
                    "Provisioned Linux/Windows servers for prod & dev; operated Nutanix clusters and VMware vSphere (vCenter 8).",
                    "Managed AWS (EC2, S3, Route53, VPC, FSx); implemented Autoscaling via Lambda + EventBridge.",
                    "Operated mail (Exchange/Zimbra) and Barracuda Email Security; enforced MFA on O365.",
                    "Ran DR: Veeam, Acronis, Nutanix protection; maintained AD Connect and Azure Entra ID.",
                    "Led migration of ~69 workloads from Nutanix AHV to Dell PowerEdge R660 on ESXi vSphere 7.",
                },
            },
            {
                Title: "IT Datacenter Specialist",
                Company: "PT. Transretail Indonesia",
                Location: "Jakarta",
                Period: "2018–2022",
                Bullets: []string{
                    "Monitored, diagnosed, and resolved complex network & server incidents.",
                    "Implemented full file-sharing backup with Veritas Backup Exec & Duplicati.",
                    "Built an incident portal using OSticket for DC operations.",
                },
            },
        },
        Certifications: []Certification{
            {Name: "VCTA-DCV 2024 (VMware)", Issuer: "VMware", Year: "2024", Note: ""},
            {Name: "AWS Academy Cloud Operations", Issuer: "AWS Academy", Year: "2023", Note: ""},
            {Name: "NSE 1 & NSE 2", Issuer: "Fortinet (expired 2022)", Year: "2020", Note: ""},
            {Name: "Alibaba Cloud (expired 2022)", Issuer: "Alibaba", Year: "2020", Note: ""},
        },
        Education: []Education{
            {Degree: "Bachelor of IT (Information Systems)", School: "Gunadarma University", Period: "2011–2015", GPA: "3.07"},
        },
        Year: time.Now().Year(),
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
