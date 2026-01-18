import { useState } from "react";
import axios from "axios";

function GetResultById() {
  const [analysisId, setAnalysisId] = useState("");
  const [data, setData] = useState<any | null>(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const fetchResult = async () => {
    if (!analysisId.trim()) {
      setError("Analysis ID is required");
      return;
    }

    try {
      setLoading(true);
      setError("");
      setData(null);

      const res = await axios.get(
        `http://localhost:8080/api/entity-classification/results/${analysisId}`
      );

      setData(res.data);
    } catch (err) {
      setError("Invalid Analysis ID or result not found");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="page">
  <h2>Get Result by Analysis ID</h2>

  <input
    type="text"
    placeholder="Enter analysis_id"
    value={analysisId}
    onChange={(e) => setAnalysisId(e.target.value)}
    style={{ width: "420px" }}
  />

  <br /><br />

  <button onClick={fetchResult} disabled={loading}>
    {loading ? "Fetching..." : "Get Result"}
  </button>

  {error && <p className="error">{error}</p>}
      {data && (
        <table
          border={1}
          cellPadding={8}
          style={{ marginTop: 20, width: "100%" }}
        >
          <thead>
            <tr>
              <th>People</th>
              <th>Tools</th>
              <th>Brands</th>
              <th>Products</th>
              <th>Tone</th>
              <th>Style</th>
              <th>Sensitive Domains</th>
              <th>Severity</th>
              <th>Requires Review</th>
              <th>Created At</th>
            </tr>
          </thead>

          <tbody>
            <tr>
              <td>{data.entities?.people?.join(", ") || "-"}</td>
              <td>{data.entities?.tools?.join(", ") || "-"}</td>
              <td>{data.entities?.brands?.join(", ") || "-"}</td>
              <td>{data.entities?.products?.join(", ") || "-"}</td>
              <td>{data.tone?.primary || "-"}</td>
              <td>{data.style?.primary || "-"}</td>
              <td>
                {data.safety_flags?.sensitive_domains?.length
                  ? data.safety_flags.sensitive_domains.join(", ")
                  : "None"}
              </td>
              <td>{data.safety_flags?.severity || "None"}</td>
              <td>
                {data.safety_flags?.requires_review ? "Yes" : "No"}
              </td>
              <td>
                {data.created_at
                  ? new Date(data.created_at).toLocaleString()
                  : "-"}
              </td>
            </tr>
          </tbody>
        </table>
      )}
    </div>
  );
}

export default GetResultById;
