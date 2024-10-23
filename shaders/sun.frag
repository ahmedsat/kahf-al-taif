#version 450 core

out vec4 FragColor;   

uniform sampler2D uTexture; 

void main() {

    FragColor = texture(uTexture, gl_FragCoord.xy);
   
    FragColor = vec4(1.0);
}