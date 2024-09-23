#version 450 core

layout(location = 0) in vec3 vp;
layout(location = 1) in vec3 color;
out vec3 frag_color;
void main() {

    frag_color = color;

    gl_Position = vec4(vp, 1.0);
}