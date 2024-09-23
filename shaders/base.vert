#version 450 core

layout(location = 0) in vec3 vp;
layout(location = 1) in vec3 color;
layout(location = 2) in vec2 texcoord;

out vec3 frag_color;
out vec2 frag_texcoord;

void main() {

    frag_color = color;
    frag_texcoord = texcoord;

    gl_Position = vec4(vp, 1.0);
}