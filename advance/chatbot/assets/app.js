document.addEventListener("DOMContentLoaded", () => {
  const chatLog = document.getElementById("chat-log");
  const chatInput = document.getElementById("chat-input");
  const sendBtn = document.getElementById("send-btn");

  sendBtn.addEventListener("click", () => {
    const message = chatInput.value.trim();
    if (!message) return;
    addMessageToChatLog(message, "user"); // ① まず画面にユーザーメッセージを表示
    chatInput.value = "";
    // ② サーバーにメッセージを送信 (POST -> SSE 受信)
    // fetchでPOST + SSE 受信の例
    sendMessageToServer(message);
  });

  function addMessageToChatLog(text, sender) {
    const newMessage = document.createElement("div");
    newMessage.classList.add("chat-message");
    if (sender === "user") {
      newMessage.classList.add("user");
    } else {
      newMessage.classList.add("bot");
    }
    newMessage.innerText = text;
    chatLog.appendChild(newMessage);
    // スクロールを最新メッセージまで下げたい時は以下
    chatLog.scrollTop = chatLog.scrollHeight;
  }

  function sendMessageToServer(message) {
    const url = new URL('/api/conversation', location.href);
    url.searchParams.append('message', message);
    /* 
      以下、SSEの簡易例（EventSource を使う場合）。
      実際にはPOSTとSSE(=GET)を分離することが多いですが、
      チュートリアルの都合上この書き方をしています。

      ※ より正確には:
        - POSTでメッセージを送る
        - 受信は GET /api/conversation のSSEをEventSourceで開く
        というように分けることが多い

      ここではあくまで概念的なサンプルにとどめます。
    */
    const eventSource = new EventSource(url);
    const messages = [];
    eventSource.onmessage = function (e) {
      // SSEで受信したメッセージをbot吹き出しで追加
      if (e.data === '[DONE]') {
        eventSource.close();
        addMessageToChatLog(messages.join(''), "bot");
      } else {
        messages.push(e.data);
      }
    };

    eventSource.onerror = function (err) {
      console.error("SSE error:", err);
      eventSource.close();
    };
  }
});
