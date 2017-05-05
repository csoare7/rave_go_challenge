var tracker = new function() {
  this.url = null;
  this.sessionId = null;

  this.init = function() {
    this.setUrl();
    this.sessionId();
  }
  
  this.setUrl = function() {
    this.url = window.location.origin;
  }

  this.getUrl = function() {
    return this.url;
  }

  this.setSessionId = function(sessionId) {
    if (typeof(Storage) !== "undefined") {
      this.sessionId = localStorage.setItem("ravelin_sessionId", sessionId);
    } else {
        // No web storage support. To implement revert back to cookie
    }
  }

  this.getSessionId = function() {
    return sessionId; // or localStorage.getItem("ravelin_sessionId")
  }

  this.addResizeListener = function() {}
  this.addCopyPasteListener = function() {}
  
}

tracker.init()