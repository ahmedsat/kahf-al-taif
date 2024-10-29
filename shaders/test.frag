#version 330 core

in vec3 vColor;
in vec2 vTexCoord;

out vec4 FragColor;

uniform vec3 uForegroundColor;
uniform sampler2D uWallTexture;

void main()
{
    FragColor = vec4(vColor, 1.0f);
    // FragColor = texture(uWallTexture, vTexCoord);
} 