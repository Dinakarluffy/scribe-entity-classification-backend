import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const analyzeTranscript = async (payload: {
  transcript_text: string;
}) => {
  return axios.post(
    `${API_BASE_URL}/api/entity-classification/analyze`,
    payload
  );
};

export const getResults = async () => {
  return axios.get(
    `${API_BASE_URL}/api/entity-classification/results`
  );
};
