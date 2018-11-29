package main

import (
    "os"
    "io/ioutil"
    "fmt"
    "debug/pe"
    . "github.com/owlinux1000/readpe/header"
    "github.com/owlinux1000/clib"
)

func main() {

    app, err := clib.NewApp("readpe", "1.0.0", "PE Parser like readelf")
    if err != nil {
        panic(err)
    }
    app.NoArgUsageFlag = true

    if err := app.AddOption("-d", "Display the MS-DOS header", 1); err != nil {
        panic(err)
    }

    if err := app.AddOption("-s", "Display section headers", 1); err != nil {
        panic(err)
    }

    if err := app.AddOption("-o", "Display the optional header", 1); err != nil {
        panic(err)
    }

    if err := app.AddOption("-f", "Display file header", 1); err != nil {
        panic(err)
    }


    exitStatus, err := app.Parse(os.Args)
    if err != nil {
        fmt.Println(err)
        os.Exit(exitStatus)
    }
    if len(os.Args) == 1 {
        os.Exit(0)
    }
    
    if ok, _ := app.OptionFlag("-h"); ok {
        fmt.Println(app.Usage())
        os.Exit(exitStatus)
    }

    if ok, _ := app.OptionFlag("-v"); ok {
        fmt.Println(app.Version)
        os.Exit(exitStatus)
    }

    if ok, _ := app.OptionFlag("-d"); ok {
        args, err := app.OptionArgs("-d")
        if err != nil {
            panic(err)
        }
        data, err := ioutil.ReadFile(args[0])
        if err != nil {
            panic(err)
        }
        h := NewImageDosHeader(data)
        fmt.Printf("%s", h)
        os.Exit(0)
    }

    path := os.Args[2]
    file, err := pe.Open(path)
    if err != nil {
        panic(err)
    }

    if ok, _ := app.OptionFlag("-o"); ok {
        switch file.OptionalHeader.(type) {
        case *pe.OptionalHeader32:
            fmt.Printf("%s", ImageOptionalHeader32(*file.OptionalHeader.(*pe.OptionalHeader32)))
        case *pe.OptionalHeader64:
            fmt.Printf("%s", ImageOptionalHeader64(*file.OptionalHeader.(*pe.OptionalHeader64)))
        }
    }

    if ok, _ := app.OptionFlag("-s"); ok {
        for _, section := range file.Sections {
            fmt.Printf("%s", ImageSectionHeader(section.SectionHeader))
        }
    }
    
    if ok, _ := app.OptionFlag("-f"); ok {
        fmt.Printf("%s", ImageFileHeader(file.FileHeader))
    }

}
