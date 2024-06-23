import streamlit as st
import requests

# Streamlit UI
st.title("AI Şiir Oluşturucu 🎭")
st.write("Şiir oluşturmak için bir prompt giriniz")

prompt = st.text_input("Prompt")

if st.button("Şiir Oluştur"):
    if prompt:
        # Make a request to the Go backend
        st.write("Prompt: ", prompt)  # Prompt'i loglamak
        response = requests.post("http://localhost:8000/generate_poem", json={"prompt": prompt})
        st.write("Response Status Code: ", response.status_code)  # Yanıt durumu loglama
        st.write("Response Content: ", response.content)  # Yanıt içeriğini loglama
        if response.status_code == 200:
            poem = response.json().get("poem")
            if poem:
                st.write("### Oluşturulan şiir:")
                st.write(poem)
            else:
                st.error("Poem generation failed, no content received.")
        else:
            st.error(f"Failed to generate poem, status code: {response.status_code}")
    else:
        st.error("Prompt cannot be empty")
