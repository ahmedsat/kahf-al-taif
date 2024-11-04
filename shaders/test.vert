#version 330 core
layout (location = 0) in vec3 aPosition;
layout (location = 1) in vec2 aTexCoord;
layout (location = 2) in vec3 aColor;
layout (location = 3) in vec3 aNormal;

out vec2 vTexCoord;
out vec3 vColor;
out vec3 vNormal;
out vec3 vFragPos;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;
uniform mat3 uNormalMatrix;



void main()
{
    vec4 pos = uProjection * uView * uModel * vec4(aPosition, 1.0);
    gl_Position = pos;
    vTexCoord = aTexCoord;
    vColor = aColor;
    vNormal = uNormalMatrix * aNormal;
    vFragPos = vec3(uModel * vec4(aPosition, 1.0));
}