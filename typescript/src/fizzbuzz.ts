export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);

  let no = 0;
  const getFizzBuzz = () => {
    ++no;
    if (no % 3 === 0 || no % 5 === 0) {
      return (no % 3 === 0 ? "Fizz" : "") + (no % 5 === 0 ? "Buzz" : "");
    }
    return no.toString();
  };

  button.addEventListener("click", () => {
    const li = document.createElement("li");
    li.innerHTML = getFizzBuzz();
    list.appendChild(li);
  });
}
