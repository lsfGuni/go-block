# Go-Block

<p><strong>Go-Block</strong> is a Go-based blockchain application designed to securely store and retrieve medicine tracking data using Ethereum smart contracts. It ensures an immutable, transparent record of medicine shipments and sales for enhanced traceability and auditability, providing a trustless system for stakeholders.</p>

<h2>Table of Contents</h2>

<ol>
  <li><a href="#project-overview">Project Overview</a></li>
  <li><a href="#requirements">Requirements</a></li>
  <li><a href="#installation">Installation</a></li>
  <li><a href="#running-the-project">Running the Project</a></li>
  <li><a href="#testing-the-project-api">Testing the Project API</a></li>
  <li><a href="#endpoints">Endpoints</a></li>
</ol>

---

<h2 id="project-overview">Project Overview</h2>

<p>The <strong>go-block</strong> application acts as an API middleware between a backend and the Ethereum blockchain. It enables:</p>
<ul>
  <li><strong>Data Integrity</strong>: Uses hashing and immutability for tamper-proof records.</li>
  <li><strong>Transparency and Auditability</strong>: Provides a complete history of each medicine’s shipment and sale records.</li>
</ul>

<h2 id="requirements">Requirements</h2>

<p>To run the <strong>go-block</strong> project, ensure you have the following installed:</p>
<ul>
  <li><strong>Go</strong>: v1.16 or higher. Download from <a href="https://golang.org/dl/">GoLang</a>.</li>
  <li><strong>Geth (Go-Ethereum)</strong>: Required to interact with an Ethereum network. Install from <a href="https://geth.ethereum.org/downloads/">Geth</a>.</li>
  <li><strong>Solidity Compiler</strong>: If deploying smart contracts, you can use <a href="https://remix.ethereum.org/">Remix IDE</a> or install <code>solc</code> locally.</li>
</ul>

---

<h2 id="installation">Installation</h2>

<h3>1. Clone the Repository</h3>

<pre><code>git clone https://github.com/lsfGuni/go-block.git
cd go-block
</code></pre>

<h3>2. Set Up Dependencies</h3>

<pre><code>go mod tidy
</code></pre>

<h3>3. Install Ethereum and Geth</h3>

<ol>
  <li><strong>Download and Install Geth</strong>:
    <ul>
      <li>Install Geth from the <a href="https://geth.ethereum.org/downloads/">official website</a>.</li>
      <li>Verify installation:
      <pre><code>geth version</code></pre></li>
    </ul>
  </li>

  <li><strong>Start a Local Ethereum Node</strong> (Optional, if you want to run a private network):
  <pre><code>geth --http --http.addr "localhost" --http.port 8545 --http.api "eth,net,web3" console</code></pre>
  </li>

  <li><strong>Deploy the Smart Contract</strong>:
    <ul>
      <li>Write and deploy your <code>MedicineTracking</code> contract using <a href="https://remix.ethereum.org/">Remix IDE</a> or use the command line with <code>solc</code> and Geth.</li>
      <li>Note down the <strong>contract address</strong> and <strong>ABI</strong>.</li>
    </ul>
  </li>
</ol>

<h3>4. Configure Go Project</h3>

<p>Edit <code>main.go</code> to include the contract address, ABI, and Ethereum node details.</p>

<pre><code>const contractAddress = "YOUR_CONTRACT_ADDRESS"
const infuraURL = "https://ropsten.infura.io/v3/YOUR_INFURA_PROJECT_ID"
</code></pre>

---

<h2 id="running-the-project">Running the Project</h2>

<ol>
  <li><strong>Run Geth</strong> (or connect to Infura if using an Ethereum testnet).</li>
  <li><strong>Run the Go Server</strong>:
    <pre><code>go run main.go</code></pre>
    <p>The server will be available on <code>http://localhost:8081</code>.</p>
  </li>
</ol>

---

<h2 id="testing-the-project-api">Testing the Project API</h2>

<p>The API can be tested using command-line tools like <code>curl</code> or API testing clients like Postman.</p>

<ol>
  <li><strong>Run the Client Application (csvfront)</strong>
    <ul>
      <li>Start the <code>csvfront</code> client on port <code>8082</code> to initiate CSV uploads.</li>
    </ul>
  </li>
  <li><strong>Run the API Server (drug-track)</strong>
    <ul>
      <li>Start the API server on port <code>8080</code> to handle CSV uploads and interact with <code>go-block</code>.</li>
    </ul>
  </li>
  <li><strong>Send Data to the Blockchain Server (go-block)</strong>
    <ul>
      <li>The <code>drug-track</code> server will call <code>go-block</code> on <code>http://localhost:8081</code> to store data on the Ethereum blockchain.</li>
    </ul>
  </li>
</ol>

---

<h2 id="endpoints">Endpoints</h2>

<h3>1. <code>POST /storeData</code></h3>

<ul>
  <li><strong>Description</strong>: Stores medicine tracking data (shipment or sales records) on the blockchain.</li>
  <li><strong>Request Body</strong>:
    <pre><code>{
  [
    {
      "seq": 14707,
      "hashCode": "37565e8afbe3a9b525211fe3f68b4862af4e4f66a776ce900563f2f265eea032",
      "status": "Shipped"
    }
  ]
}</code></pre>
  </li>
  <li><strong>Response</strong>:
    <pre><code>{
  "status": "success"
}</code></pre>
  </li>
</ul>

<h3>2. <code>GET /getRecords</code></h3>

<ul>
  <li><strong>Description</strong>: Retrieves all stored records or a specific record by <code>seq</code>.</li>
  <li><strong>Query Parameters</strong>:
    <ul>
      <li><code>seq</code> (optional): Unique identifier of the record.</li>
    </ul>
  </li>
  <li><strong>Example</strong>:
    <pre><code>curl "http://localhost:8081/getRecords?seq=14707"</code></pre>
  </li>
  <li><strong>Response</strong>:
    <pre><code>{
  "seq": 14707,
  "hashCode": "37565e8afbe3a9b525211fe3f68b4862af4e4f66a776ce900563f2f265eea032",
  "status": "Shipped",
  "timestamp": 1637924800
}</code></pre>
  </li>
</ul>
