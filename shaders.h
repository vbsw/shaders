#ifndef SHADERS_H
#define SHADERS_H

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

typedef const char* constchar_t;

extern const char *const primitive_vertex_shader;
extern const char *const primitive_fragment_shader;
extern const char *const primitive_position_attribute;
extern const char *const primitive_color_attribute;
extern const char *const primitive_projection_uniform;
extern const char *const primitive_model_uniform;

#ifdef __cplusplus
}
#endif // __cplusplus

#endif /* SHADERS_H */