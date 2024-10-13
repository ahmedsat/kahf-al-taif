#version 450 core

in vec4 vColor;       
in vec2 vTexCoord;    

out vec4 FragColor;   

uniform sampler2D uTexture; 

void main() {
   
    vec4 texColor = texture(uTexture, vTexCoord);
   
    FragColor = texColor/* * vColor */;
}
