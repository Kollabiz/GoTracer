package ObjParser

import (
	"RayTracer/Maths"
	"RayTracer/RayTracing/Structures"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseObjFile(fileName string) (Meshes []Structures.Mesh, Materials []Structures.Material) {
	var meshes []Structures.Mesh
	var processingMesh *Structures.Mesh
	var currentLib []Structures.Material
	var currentMaterial *Structures.Material
	var vertices []Structures.Vertex
	var texCoords []Maths.Vector2
	var normals []Maths.Vector3
	flData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(flData), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineParts := strings.Split(line, " ")
		switch lineParts[0] {
		case "mtllib":
			currentLib = ParseMtlFile(objPathToMtlPath(fileName, lineParts[1]))
		case "usemtl":
			mat := findMaterial(currentLib, lineParts[1])
			if mat == nil {
				panic(fmt.Sprintf("unable to find material %s", lineParts[1]))
			}
			currentMaterial = mat
		case "o":
			meshes = append(meshes, Structures.MakeMesh(lineParts[1]))
			processingMesh = &meshes[len(meshes)-1]
		case "v":
			values := ParseFloats(lineParts)
			vertices = append(vertices, Structures.Vertex{Position: Maths.MakeVector3(values[0], values[1], values[2])})
		case "vt":
			values := ParseFloats(lineParts)
			texCoords = append(texCoords, Maths.Vector2{X: values[0], Y: values[1]})
		case "vn":
			values := ParseFloats(lineParts)
			normals = append(normals, Maths.Vector3{X: values[0], Y: values[1], Z: values[2]})
		case "f":
			var vertexIndices [3]int
			var texCoordIndices [3]int
			var normalIndices [3]int
			for j := 0; j < 3; j++ {
				indices := strings.Split(lineParts[j+1], "/")
				vertexIndices[j], _ = strconv.Atoi(indices[0])
				if len(indices) > 1 && indices[1] != "" {
					texCoordIndices[j], _ = strconv.Atoi(indices[1])
				}
				if len(indices) > 2 {
					normalIndices[j], _ = strconv.Atoi(indices[2])
				}
			}
			v1 := &vertices[vertexIndices[0]-1]
			v2 := &vertices[vertexIndices[1]-1]
			v3 := &vertices[vertexIndices[2]-1]

			v1.UV = texCoords[texCoordIndices[0]-1]
			v2.UV = texCoords[texCoordIndices[1]-1]
			v3.UV = texCoords[texCoordIndices[2]-1]

			v1.VertexNormal = normals[normalIndices[0]-1]
			v2.VertexNormal = normals[normalIndices[1]-1]
			v3.VertexNormal = normals[normalIndices[2]-1]

			tri := Structures.MakeTriangle(v1, v2, v3, currentMaterial)
			processingMesh.Triangles = append(processingMesh.Triangles, tri)
		}
	}

	return meshes, currentLib
}

func findMaterial(materials []Structures.Material, materialName string) *Structures.Material {
	for i := 0; i < len(materials); i++ {
		if materials[i].MaterialName == materialName {
			return &materials[i]
		}
	}
	return nil
}

func objPathToMtlPath(objPath string, mtlFileName string) string {
	p := strings.Split(objPath, "\\")
	p[len(p)-1] = mtlFileName
	return strings.Join(p, "\\")
}
