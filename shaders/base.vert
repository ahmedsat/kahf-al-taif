#version 450 core

layout(location = 0) in vec3 vp;
layout(location = 1) in vec3 color;
layout(location = 2) in vec2 texcoord;

out vec3 frag_color;
out vec2 frag_texcoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {

    frag_color = color;
    frag_texcoord = texcoord;

    gl_Position = projection * view * model * vec4(vp, 1.0);

}