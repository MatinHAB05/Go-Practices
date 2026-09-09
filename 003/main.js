document.querySelectorAll(".box").forEach((box) => {
  const editIcon = box.querySelector(".edit-icon");
  const portNumber = box.querySelector(".port-number");

  editIcon.addEventListener("click", () => {
    const existingInput = box.querySelector(".edit-input");
    const existingSubmitBtn = box.querySelector(".save-button");
    let x = portNumber.textContent;
    
    if (!existingInput && !existingSubmitBtn) {
      
      portNumber.style.display = "none";
      

      
      const input = document.createElement("input");
      input.type = "number";
      input.value = portNumber.textContent;
      input.classList.add("edit-input");

      
      const submitBtn = document.createElement("button");
      submitBtn.textContent = "ثبت";
      submitBtn.classList.add("save-button"); 

      
      box.appendChild(input);
      box.appendChild(submitBtn);

      
      submitBtn.addEventListener("click", () => {
        
        if (input.value == "") {
          portNumber.textContent = x;
        } else {
          portNumber.textContent = input.value;
        }

        
        input.remove();
        submitBtn.remove();
        portNumber.style.display = "inline";
        editIcon.style.display = "inline"; 
      });

      
      input.addEventListener("keypress", (event) => {
        if (event.key === "Enter") {
          submitBtn.click(); 
        }
      });
    }
  });
});
