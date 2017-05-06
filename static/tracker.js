// singleton def
var Tracker = new function() {
  // properties
  var _self = this;
  var url = null;
  var screen = {
    width: Screen.width,
    height: Screen.height
  }

  // public initialisation method
  _self.init = function() {
    console.log(this)
    setUrl(window.location.origin);
    addResizeListener();
    addCopyPasteListener("input");
  };
  
  // private gets/sets
  var setUrl = function(currentUrl) {
    // check url and throw
    url = currentUrl;
  };

  var getUrl = function() {
    return url;
  };

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
    inputList = document.querySelectorAll("input");
    inputList.forEach(function(input) {
      input.addEventListener("paste", handlePaste);
    });
  };

  var addResizeListener = function() {
    window.addEventListener("resize", handleResize);
  };

  // event handlers
  var handlePaste = function(event) {
    fieldId = event.target.id;
    // postData
  };

  var handleResize = function(event) {
    // only one resize occurs
    window.removeEventListener("resize", handleResize);
    var newScreen = {
      width: Screen.width,
      height: Screen.height
    }
    // postData
  };

};


$(document).ready(function() {
  Tracker.init();
});