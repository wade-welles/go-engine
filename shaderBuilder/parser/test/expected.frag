#version 330

float import1() {
	return 1.0;
}

vec4 vertFunc() {
  return vec4(1.0);
}

vec4 fragFunc() {
  return vec4(0.5);
}

void main() {
  gl_FragColor = fragFunc();
}