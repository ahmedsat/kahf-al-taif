#version 330 core
in vec3 vColor;
out vec4 FragColor;

void main()
{
    FragColor = vec4(vColor.r, vColor.g, vColor.b, 1.0f);
} 