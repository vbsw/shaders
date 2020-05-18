/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package shaders

// #include <string.h>
// #include "shaders.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// Shader is for storing shaders and their attributes.
type Shader struct {
	VertexShader       **uint8
	FragmentShader     **uint8
	VertexShaderID     uint32
	FragmentShaderID   uint32
	ProgramID          uint32
	PositionAttribute  *uint8
	ColorAttribute     *uint8
	ProjectionUniform  *uint8
	ModelUniform       *uint8
	PositionLocation   int32
	ColorLocation      int32
	ProjectionLocation int32
	ModelLocation      int32
}

// VertexShaderStr returns vertex shader as string.
func (shader *Shader) VertexShaderStr() string {
	return NewString(*shader.VertexShader)
}

// FragmentShaderStr returns fragment shader as string.
func (shader *Shader) FragmentShaderStr() string {
	return NewString(*shader.FragmentShader)
}

// NewPrimitiveShader returns an instance of Shader for primitive types.
//
// VertexShader contains the following program:
//
//	#version 130
//
//	in vec3 positionIn;
//	in vec4 colorIn;
//	out vec4 fragementColor;
//
//	uniform mat4 projection = mat4(1.0);
//	uniform mat4 model = mat4(1.0);
//
//	void main() {
//		gl_Position = projection * model * vec4(positionIn, 1.0f);
//		fragementColor = colorIn;
//	}
//
// FragmentShader contains the following program:
//
//	#version 130
//
//	in vec4 fragementColor;
//	out vec4 color;
//
//	void main () {
//		color = fragementColor;
//	}
func NewPrimitiveShader() *Shader {
	shader := new(Shader)
	shader.VertexShader = (**uint8)(unsafe.Pointer(&C.primitive_vertex_shader))
	shader.FragmentShader = (**uint8)(unsafe.Pointer(&C.primitive_fragment_shader))
	shader.PositionAttribute = (*uint8)(unsafe.Pointer(C.primitive_position_attribute))
	shader.ColorAttribute = (*uint8)(unsafe.Pointer(C.primitive_color_attribute))
	shader.ProjectionUniform = (*uint8)(unsafe.Pointer(C.primitive_projection_uniform))
	shader.ModelUniform = (*uint8)(unsafe.Pointer(C.primitive_model_uniform))
	return shader
}

// NewString converts c-string to string and returns it. Data of c-string is not reallocated,
// but used in string.
func NewString(cstr *uint8) string {
	ccstr := C.constchar_t(unsafe.Pointer(cstr))
	n := int(C.strlen(ccstr))
	dataSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ccstr)),
		Len:  n,
		Cap:  n,
	}))
	str := string(dataSlice)
	return str
}
