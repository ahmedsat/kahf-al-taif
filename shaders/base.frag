#version 450 core

in vec2 vTexCoord;    
in vec3 vNormal;
in vec3 vFragPos;
in vec3 vViewPos;

out vec4 FragColor;   

uniform sampler2D uTexture; 

uniform vec3 uAmbientColor; 
uniform float uAmbientStrength;


void main() {
    vec4 baseColor = texture(uTexture, vTexCoord);

    vec3 ambient = uAmbientColor * uAmbientStrength;

    vec3 finalColor = ambient * baseColor.rgb;

    FragColor = vec4(finalColor, baseColor.a);
}
