package ObjParser

import (
	"RayTracer/Color"
	"RayTracer/RayTracing/Structures"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseMtlFile(fileName string) []Structures.Material {
	var parsedMaterials []Structures.Material
	var currentMaterial *Structures.Material
	flData, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("unable to open mtl file %s", fileName))
	}
	splitLines := strings.Split(string(flData), "\n")
	for lineNo := 0; lineNo < len(splitLines); lineNo++ {
		line := splitLines[lineNo]
		lineParts := strings.Split(line, " ")
		switch lineParts[0] {
		case "newmtl":
			parsedMaterials = append(parsedMaterials, *Structures.NewMaterial(lineParts[1]))
			currentMaterial = &parsedMaterials[len(parsedMaterials)-1]
		case "Kd": // Diffuse color
			values := ParseFloats(lineParts)
			currentMaterial.DiffuseColor = Color.Color{R: values[0], G: values[1], B: values[2]}
		case "Ks": // Specular color
			values := ParseFloats(lineParts)
			currentMaterial.SpecularColor = Color.Color{R: values[0], G: values[1], B: values[2]}
		case "Ke": // Emission color
			values := ParseFloats(lineParts)
			currentMaterial.EmissionColor = Color.Color{R: values[0], G: values[1], B: values[2]}
		case "Ns": // Specular exponent
			value := ParseFloats(lineParts)[0]
			currentMaterial.Glossiness = value / 2000
		case "map_Kd": // Diffuse texture
			texture := Structures.Texture2FromFile(lineParts[1])
			currentMaterial.DiffuseTexture = texture
		case "map_Ks": // Specular texture
			texture := Structures.Texture2FromFile(lineParts[1])
			currentMaterial.SpecularTexture = texture
		case "map_Ke": // Emission texture
			texture := Structures.Texture2FromFile(lineParts[1])
			currentMaterial.EmissionTexture = texture
		}
	}
	return parsedMaterials
}

func ParseFloats(data []string) []float32 {
	var parsed []float32
	for i := 1; i < len(data); i++ {
		conv, err := strconv.ParseFloat(data[i], 32)
		if err != nil {
			panic(fmt.Sprintf("invalid float literal %s", data[i]))
		}
		parsed = append(parsed, float32(conv))
	}
	return parsed
}
