export default {
  async fetch(request, env) {
    const url = new URL(request.url);
    const cmd = url.searchParams.get('cmd') || 'help';
    const responses = {
      help: 'Available commands: help, skills, projects, stats, whoami, uptime',
      skills: 'Go 95% | Python 82% | Docker 78% | Security 85%',
      projects: '1. ShieldGuard 2. AetherLang 3. AetherSentinel',
      stats: 'Stars: 150+ | Repos: 106 | Org: AetherCode-Core',
      whoami: 'AetherCodeHQ - Core Developer & Maintainer'
    };
    return new Response(JSON.stringify({ command: cmd, output: responses[cmd] || 'Unknown command' }), {
      headers: { 'Content-Type': 'application/json' }
    });
  }
};
