#version 330 core
layout (location = 0) in vec3 aPosition;
layout (location = 1) in vec2 aTexCoord;
layout (location = 2) in vec3 aColor;

out vec2 vTexCoord;
out vec3 vColor;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;



void main()
{
    vec4 pos = uProjection * uView * uModel * vec4(aPosition, 1.0);
    gl_Position = pos;
    vTexCoord = aTexCoord;
    vColor = aColor;
}