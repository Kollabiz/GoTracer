package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
)

// Material represents a material for meshes, containing diffuse, specular, emission and normal information
type Material struct {
	DiffuseColor     Color.Color
	SpecularColor    Color.Color
	EmissionColor    Color.Color
	EmissionStrength float32
	Glossiness       float32
	// Textures
	DiffuseTexture  *Texture2
	SpecularTexture *Texture2
	EmissionTexture *Texture2
	NormalMap       *Texture2
	MaterialName    string
}

func NewMaterial(materialName string) *Material {
	mat := new(Material)
	mat.MaterialName = materialName
	return mat
}

func (mat *Material) GetDiffuse(uv Maths.Vector2) Color.Color {
	diff := mat.DiffuseColor
	if mat.DiffuseTexture != nil {
		diff.IMul(mat.DiffuseTexture.GetPixelFromUV(uv))
	}
	return diff
}

func (mat *Material) GetSpecular(uv Maths.Vector2) Color.Color {
	spec := mat.SpecularColor
	if mat.SpecularTexture != nil {
		spec.IMul(mat.SpecularTexture.GetPixelFromUV(uv))
	}
	return spec
}

func (mat *Material) GetEmission(uv Maths.Vector2) Color.Color {
	emm := mat.SpecularColor
	if mat.SpecularTexture != nil {
		emm.IMul(mat.SpecularTexture.GetPixelFromUV(uv))
	}
	return emm
}

func (mat *Material) GetNormal(uv Maths.Vector2) Maths.Vector3 {
	clr := mat.NormalMap.GetPixelFromUV(uv)
	norm := Maths.MakeVector3(clr.R, clr.G, clr.B).Normalized()
	return norm
}
