#version 450 core

out vec4 FragColor;

in vec2 frag_texcoord;

uniform sampler2D wall;
uniform sampler2D stone;
uniform sampler2D tennant;

void main() {
    FragColor = texture(wall, frag_texcoord);
}

