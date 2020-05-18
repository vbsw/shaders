/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

#include <string.h>

const char *const primitive_vertex_shader = "#version 130\n\nin vec3 positionIn;\nin vec4 colorIn;\nout vec4 fragementColor;\n\nuniform mat4 projection = mat4(1.0);\nuniform mat4 model = mat4(1.0);\n\nvoid main() {\n\tgl_Position = projection * model * vec4(positionIn, 1.0f);\n\tfragementColor = colorIn;\n}";
const char *const primitive_fragment_shader = "#version 130\n\nin vec4 fragementColor;\nout vec4 color;\n\nvoid main() {\n\tcolor = fragementColor;\n}";
const char *const primitive_position_attribute = "positionIn";
const char *const primitive_color_attribute = "colorIn";
const char *const primitive_projection_uniform = "projection";
const char *const primitive_model_uniform = "model";
