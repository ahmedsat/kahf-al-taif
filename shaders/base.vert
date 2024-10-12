#version 450 core

layout(location = 0) in vec4 aPosition;   
layout(location = 1) in vec4 aColor;      
layout(location = 2) in vec2 aTexCoord;   

out vec4 vColor;        
out vec2 vTexCoord;     

uniform mat4 uModel;    
uniform mat4 uView;     
uniform mat4 uProjection; 

void main() {
    
    gl_Position = uProjection * uView * uModel * aPosition;
    
    vColor = aColor;
    vTexCoord = aTexCoord;
}
