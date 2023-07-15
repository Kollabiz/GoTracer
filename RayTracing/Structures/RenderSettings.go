package Structures

type RenderSettings struct {
	UseIndirectIllumination              bool
	OptimizeIndirectIlluminationRayCount bool // The more glossiness, the fewer rays are being cast
	IndirectIlluminationSampleCount      int
	IndirectIlluminationDepth            int
	DumpDebugPasses                      bool
	ProgressiveRenderingPassQuantity     int
	UseTiling                            bool
	TileWidth                            int
	TileHeight                           int
}
