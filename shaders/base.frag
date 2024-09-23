#version 450 core

out vec4 FragColor;

in vec3 frag_color;
in vec2 frag_texcoord;

uniform sampler2D wall;
uniform sampler2D stone;

void main() {
    FragColor = texture(wall, frag_texcoord) * texture(stone, frag_texcoord) * vec4(frag_color, 1.0);
}

