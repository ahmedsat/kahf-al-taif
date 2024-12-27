#version 450 core

// Input vertex attributes
in vec2 vTexCoord;
in vec3 vNormal;
in vec3 vFragPos;

// Output fragment color
out vec4 FragColor;

// Light structure for multiple light support
struct Light {
    vec3 position;
    vec3 color;
    float ambient;
    float diffuse;
    float specular;
};

// Material properties structure
struct Material {
    sampler2D diffuseMap;
    sampler2D specularMap;
    float shininess;
};

// Uniform variables
uniform vec3 uCameraPosition;
uniform Light uMainLight;
uniform Material uMaterial;

// Optional: Multiple light support (up to 4 lights)
#define MAX_LIGHTS 4
uniform Light uAdditionalLights[MAX_LIGHTS];
uniform int uActiveLightCount;

// Optional: Fog effect
uniform bool uEnableFog;
uniform vec3 uFogColor;
uniform float uFogDensity;

vec3 calculateLighting(Light light, vec3 normal, vec3 fragPos, vec3 viewDir) {
    // Ambient component
    vec3 ambient = light.color * light.ambient * vec3(texture(uMaterial.diffuseMap, vTexCoord));

    // Diffuse component
    vec3 lightDir = normalize(light.position - fragPos);
    float diff = max(dot(normal, lightDir), 0.0);
    vec3 diffuse = light.color * light.diffuse * diff * vec3(texture(uMaterial.diffuseMap, vTexCoord));

    // Specular component
    vec3 reflectDir = reflect(-lightDir, normal);
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), uMaterial.shininess);
    vec3 specular = light.color * light.specular * spec * vec3(texture(uMaterial.specularMap, vTexCoord));

    return ambient + diffuse + specular;
}

void main() {
    // Normalize input vectors
    vec3 normal = normalize(vNormal);
    vec3 viewDir = normalize(uCameraPosition - vFragPos);

    // Calculate main light contribution
    vec3 lightColor = calculateLighting(uMainLight, normal, vFragPos, viewDir);

    // Optional: Add additional light contributions
    for (int i = 0; i < uActiveLightCount; i++) {
        lightColor += calculateLighting(uAdditionalLights[i], normal, vFragPos, viewDir);
    }

    vec3 finalColor = lightColor;

    // Optional: Fog effect
    if (uEnableFog) {
        float distance = length(uCameraPosition - vFragPos);
        float fogFactor = exp(-uFogDensity * distance);
        finalColor = mix(uFogColor, finalColor, fogFactor);
    }

    // Final fragment color
    FragColor = vec4(finalColor, 1.0);
}