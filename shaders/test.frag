#version 330 core

in vec3 vColor;
in vec2 vTexCoord;
in vec3 vNormal;
in vec3 vFragPos;

out vec4 FragColor;

uniform sampler2D uWallTexture;

uniform float uAmbientStrength;
uniform vec3 uAmbientColor;

uniform vec3 uDiffuseLightPosition;
uniform vec3 uDiffuseLightColor;


void main()
{

    vec3 ambientColor = uAmbientColor * uAmbientStrength;

    vec3 normal = normalize(vNormal);
    vec3 lightDir = normalize(uDiffuseLightPosition - vFragPos);
    float diff = max(dot(normal, lightDir), 0.0); 
    vec3 diffuse = uDiffuseLightColor * diff;

    vec3 LightColor = diffuse*diff+ambientColor*uAmbientStrength/(diff+uAmbientStrength);

    FragColor = vec4(vColor*LightColor, 1.0f);
    // FragColor = texture(uWallTexture, vTexCoord);
} 