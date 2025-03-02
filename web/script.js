document.addEventListener('DOMContentLoaded', function () {
  const form = document.getElementById('calcForm');
  const expressionInput = document.getElementById('expressionInput');
  const resultArea = document.getElementById('resultArea');

  form.addEventListener('submit', function (e) {
    e.preventDefault();
    const expression = expressionInput.value.trim();
    if (!expression) {
      alert("Пожалуйста, введите выражение.");
      return;
    }
    resultArea.innerHTML = "Отправка запроса на сервер...";

    fetch('http://localhost:8080/api/v1/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ expression })
    })
    .then(response => {
      if (!response.ok) {
        throw new Error("Ошибка сервера: " + response.status);
      }
      return response.json();
    })
    .then(data => {
      const exprId = data.id;
      resultArea.innerHTML = `Выражение отправлено. ID: ${exprId}. Ожидание результата...`;
      const pollInterval = setInterval(() => {
        fetch(`http://localhost:8080/api/v1/expressions/${exprId}`)
        .then(response => {
          if (!response.ok) {
            throw new Error("Ошибка получения результата: " + response.status);
          }
          return response.json();
        })
        .then(exprData => {
          if (exprData.status === "done") {
            resultArea.innerHTML = `Результат для выражения "${exprData.expression}" = ${exprData.result}`;
            clearInterval(pollInterval);
          } else {
            resultArea.innerHTML = `Статус: ${exprData.status}. Ожидание завершения вычислений...`;
          }
        })
        .catch(error => {
          console.error("Ошибка опроса результата:", error);
          resultArea.innerHTML = "Ошибка получения результата.";
          clearInterval(pollInterval);
        });
      }, 2000);
    })
    .catch(error => {
      console.error("Ошибка отправки выражения:", error);
      resultArea.innerHTML = "Ошибка отправки выражения.";
    });
  });
});
