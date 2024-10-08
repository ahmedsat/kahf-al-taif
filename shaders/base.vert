#version 450 core

layout(location = 0) in vec3 vp;
layout(location = 1) in vec2 texcoord;

out vec2 frag_texcoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {

    frag_texcoord = texcoord;

    gl_Position =  projection *  view * model * vec4(vp, 1.0);

}