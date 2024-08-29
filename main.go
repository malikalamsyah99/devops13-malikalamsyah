package main

import (
    "html/template"
    "net/http"
)

type CV struct {
    Name      string
    Position  string
    Summary   string
    Skills    []string
    Experience []Experience
    Education []Education
}

type Experience struct {
    Company  string
    Role     string
    Duration string
    Details  []string
}

type Education struct {
    Institution string
    Degree      string
    Year        string
}

func cvHandler(w http.ResponseWriter, r *http.Request) {
    // Data untuk CV
    cv := CV{
        Name:     "Malik Alamsyah",
        Position: "Software Engineer",
        Summary:  "An experienced software engineer with a strong background in Golang and web development.",
        Skills: []string{
            "Golang",
            "Docker",
            "Linux",
            "REST API",
        },
        Experience: []Experience{
            {
                Company:  "Tech Solutions Ltd.",
                Role:     "Backend Developer",
                Duration: "2019 - Present",
                Details: []string{
                    "Developed microservices using Golang.",
                    "Integrated RESTful APIs with frontend applications.",
                    "Worked on performance optimization and scalability.",
                },
            },
            {
                Company:  "Web Innovators Inc.",
                Role:     "Junior Developer",
                Duration: "2017 - 2019",
                Details: []string{
                    "Assisted in the development of web applications using PHP and JavaScript.",
                    "Maintained and updated existing web projects.",
                    "Collaborated with senior developers on various tasks.",
                },
            },
        },
        Education: []Education{
            {
                Institution: "ABC University",
                Degree:      "Bachelor of Computer Science",
                Year:        "2013 - 2017",
            },
            {
                Institution: "XYZ Vocational High School",
                Degree:      "Diploma in Software Engineering",
                Year:        "2010 - 2013",
            },
        },
    }

    // Parsing template
    tmpl := template.Must(template.ParseFiles("cv.html"))

    // Rendering template dengan data CV
    tmpl.Execute(w, cv)
}

func main() {
    http.HandleFunc("/", cvHandler)
    http.ListenAndServe(":8080", nil)
}
