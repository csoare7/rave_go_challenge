// singleton def
var Tracker = new function() {
  // properties
  var _self = this;
  var url = window.location.origin;
  var apiUrl = "http://localhost:3000/data";
  var screen = {
    width: window.innerWidth,
    height: window.innerHeight
  };
  var timeTaken = 0;
  var resizeId;

  // public initialisation method
  _self.init = function() {
    addResizeListener();
    addCopyPasteListener("input");
    addKeydownListener("input");
    addOnSubmitListener("form-details");
  };
  
  // private gets/sets
  var setSessionId = function(sessionId) {
    if (typeof(Storage) !== "undefined") {
      localStorage.setItem("ravelin_sessionId", sessionId);
    } else {
        // No web storage support. To implement revert back to cookie
    }
  };

  var getSessionId = function() {
    return localStorage.getItem("ravelin_sessionId");
  };

  // private event listeners
  var addCopyPasteListener = function(selector) {
    inputList = document.querySelectorAll(selector);
    inputList.forEach(function(input) {
      input.addEventListener("paste", pasteHandler);
    });
  };

  var addKeydownListener = function(selector) {
    inputList = document.querySelectorAll(selector);
    inputList.forEach(function(input) {
      input.addEventListener("keydown", keyDownHandler);
    });
  };

  var addResizeListener = function() {
    window.addEventListener("resize", resizeHandler);
  };

  var addOnSubmitListener = function(selector) {
    var elements = document.getElementsByClassName(selector);
    if (elements.length > 0) {
      var form = elements[0];
      form.addEventListener("submit", submitHandler);
    }
  };

  var removeResizeListener = function() {
    window.removeEventListener("resize", resizeHandler);
  };

  // event handlers
  var pasteHandler = function(event) {
    fieldId = event.target.id;
    var data = {
      "eventType": "copyAndPaste",
      "websiteUrl": url,
      "sessionId": getSessionId(),
      "fieldId": fieldId,
      "pasted": true
    }
    postData(data);
  };

  var keyDownHandler = function(event) {
    timeTaken = timeTaken || Date.now();
  }

  var resizeHandler = function(event) {
    if (resizeId !== undefined) {
      clearTimeout(resizeId);
    }

    resizeId = setTimeout(function() {
      removeResizeListener();
      var newWidth = window.innerWidth;
      var newHeight =  window.innerHeight;
      var data = {
        "eventType": "resize",
        "websiteUrl": url,
        "sessionId": getSessionId(),
        "resizeFromWidth": screen.width.toString(),
        "resizeFromHeight": screen.height.toString(),
        "resizeToWidth": newWidth.toString(),
        "resizeToHeight": newHeight.toString()
      }
      postData(data);
      // set current w/h to new variables
      screen.width = newWidth;
      screen.height = newHeight;
    }, 500);
  };


  var submitHandler = function(event) {
    event.preventDefault();
    var totalTime = miliToSec(Date.now() - timeTaken);
    var data = {
      "eventType": "timeTaken",
      "websiteUrl": url,
      "sessionId": getSessionId(),
      "time": totalTime
    };
    postData(data);
  };

  var onSuccessHandler = function(data, textStatus, xhr) {
    // set session
    if (textStatus === "success" && data.hasOwnProperty("sessionId") && data.sessionId !== undefined) {
      setSessionId(data.sessionId)
    }
  };

  var onErrorHandler = function(data, textStatus, xhr) {
    console.log("Error: ", textStatus);
  };

  // helpers
  var postData = function(data) {
    console.log(data);
    $.ajax({
      type: "POST",
      url: apiUrl,
      data: JSON.stringify(data),
      dataType: "json",
      contentType: "application/json"
    })
    .done(onSuccessHandler)
    .fail(onErrorHandler)
  };

  var miliToSec = function(time) {
    return Math.round(time / 1000);
  };
};

$(document).ready(function() {
  Tracker.init();
});
